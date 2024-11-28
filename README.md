# sops-check

[![Build Status](https://github.com/Bonial-International-GmbH/sops-check/actions/workflows/ci.yml/badge.svg)](https://github.com/Bonial-International-GmbH/sops-check/actions/workflows/ci.yml)

> [!NOTE]
> This project is still in an early development stage and a lot of the desired
> features are not implemented yet.

Check SOPS files for correct and compliant usage without decrypting them to
ensure that all SOPS files are configured in the desired fashion. The goal is
to provide a security linter that safeguards the security of the data protected
by the SOPS files against common mistakes and against malicious configurations.

We are following a design-first approach, please take a look at [the design
document](docs/design.md). We are happy to hear your thoughts about it.

## Installation

The simplest way is to install the latest version via:

```sh
go install github.com/Bonial-International-GmbH/sops-check@latest
```

Finally, consult the help for usage instructions:

```sh
sops-check --help
```

## Development

Run the tests:

```sh
make coverage
```

Lint the codebase:

```sh
make lint
```

Build locally:

```sh
make build
```
