# Go Shipyard

A production-ready Go starter kit with Docker, PostgreSQL, CI, and a structured development workflow.

## Overview

This repository provides a solid foundation for building and deploying Go applications. It includes:

- Go backend with structured project layout
- PostgreSQL with migrations
- Docker-based development and production setup
- GitHub Actions for CI
- Webpack + Tailwind frontend (optional)

## Getting Started

Clone the repository and initialize the project:

```bash
git clone <your-repo>
cd <your-repo>

./init.sh \
  --module=github.com/your-org/your-app \
  --app-slug=yourapp \
  --image-repo=ghcr.io/your-org/your-app
```

Start the development environment:

```bash
make dev
```

Run migrations manually:

```bash
make migrate-up
```

## Development

The development environment is fully containerized.

Services include:
- application
- postgres
- migrations
- frontend (webpack)

Common commands:

```bash
make dev               # start all services
make connect-db        # open psql shell
make db-reset          # reset database
make test              # run full test flow
```

## Database

- PostgreSQL is used as the primary database
- Migrations are managed via `sql-migrate`
- Application tables live in a dedicated schema
- Migration state is tracked separately

## CI

The CI pipeline:

- starts PostgreSQL
- applies migrations
- runs linting and tests
- builds frontend assets

## Deployment

Deployment is handled via Docker and GitHub Actions.

Recommended workflow:

- push to `main` → run CI
- create tag (e.g. `v1.0.0`) → deploy
- optional manual deploy via workflow dispatch

## Environment

Environment variables are managed via `.env` files.

```
.env.example        # template for local development
deploy/env/.env     # production environment
```

## Notes

- The repository is designed as a template and uses placeholders
- `init.sh` replaces all placeholders with your project values
- Docker images are versioned via commit SHA or tags
