# Anduin

Anduin is an enterprise management system, currently initialized as a minimal Go project.

## Requirements

- Go 1.16 (current development baseline)

## Getting started

```bash
go run .
```

Expected output:

```text
Anduin
```

## Development

```bash
make fmt
make vet
make test
make build
```

No HTTP APIs, database integrations, or third-party dependencies are included yet.

> Go 1.16 is no longer supported with security updates. Upgrade to a supported Go version before production deployment.
