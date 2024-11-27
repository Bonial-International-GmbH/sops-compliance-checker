// Package config provides utility functions for the sops-check.
package config

import (
	"fmt"
	"io"
	"os"

	"github.com/goccy/go-yaml"
)

// Config represents the configuration for the sops-check.
type Config struct {
	AllowUnmatched bool   `json:"allowUnmatched"`
	Rules          []Rule `json:"rules"`
}

// Rule represents a single rule in the configuration.
type Rule struct {
	AllOf       []Rule `json:"allOf,omitempty"`
	AnyOf       []Rule `json:"anyOf,omitempty"`
	Match       string `json:"match,omitempty"`
	MatchRegex  string `json:"matchRegex,omitempty"`
	Not         *Rule  `json:"not,omitempty"`
	OneOf       []Rule `json:"oneOf,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
}

// Load loads the configuration from a JSON file.
func Load(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return LoadReader(file)
}

// LoadReader loads the configuration from an io.Reader.
func LoadReader(reader io.Reader) (*Config, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	if err := Validate(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// Validate validates a configuration.
func Validate(config *Config) error {
	for _, singleRule := range config.Rules {
		if err := ValidateRule(&singleRule); err != nil {
			return err
		}
	}

	return nil
}

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}

// ValidateRule validates a single rule.
func ValidateRule(rule *Rule) error {
	matchConditions := (bool2int(rule.Match != "") +
		bool2int(rule.MatchRegex != "") +
		bool2int(rule.Not != nil) +
		bool2int(len(rule.AllOf) > 0) +
		bool2int(len(rule.AnyOf) > 0) +
		bool2int(len(rule.OneOf) > 0))

	if matchConditions != 1 {
		return fmt.Errorf("Rule must exactly one match condition, got %d", matchConditions)
	}

	nestedRules := [][]Rule{
		rule.AllOf,
		rule.AnyOf,
		rule.OneOf,
	}

	if rule.Not != nil {
		if err := ValidateRule(rule.Not); err != nil {
			return err
		}
	}

	for _, rules := range nestedRules {
		for _, subRule := range rules {
			if err := ValidateRule(&subRule); err != nil {
				return err
			}
		}
	}

	return nil
}
