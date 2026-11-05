# Anduin

Anduin is an enterprise management system built with Go.

## Development baseline

- Go 1.16 (legacy development baseline)
- Gin, GORM, Viper, Uber Zap
- MySQL 8 (optional; disabled by default)

## Run

```bash
go mod tidy
go run .
```

The server starts on port 8080. No business API routes are registered yet, so requests to unimplemented paths return 404.

## Configuration

Default values are in `configs/config.yaml`. Viper supports environment variable overrides with the `ANDUIN_` prefix, for example `ANDUIN_SERVER_PORT=9090`.

To connect MySQL, set `ANDUIN_DATABASE_ENABLED=true` and `ANDUIN_DATABASE_DSN` (see `.env.example`). Environment variables must be exported by the shell or process manager; the application does not automatically load `.env`.

## Development

```bash
go mod tidy
make fmt
make vet
make test
make build
```

No business APIs, schema migrations, or authentication are implemented yet.

> Go 1.16 is end-of-life. Upgrade to a supported Go version before production deployment.
