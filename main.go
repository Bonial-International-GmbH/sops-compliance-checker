// Package main provides the entrypoint for the sops-check executable.
package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	fmt.Println("hello world")
}
