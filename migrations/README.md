# Migrations

This document outlines the development workflow when working with the databases
for each service in our project.

## Prerequisites

Ensure you have the following installed on your system:

- `postgres:17-alpine` Docker Image
- `migrate` (Database migration tool)
- DataGrip or similar Visual Database Management tool

### Environment Variables

Ensure you set each of the necessary environment variables before running
these commands.

**Example:**

```bash
export SHIPMENT_DB_CONNECTION_STRING=postgres://postgres:postgres@localhost:5432/shipment_db?sslmode=disable
export INVENTORY_DB_CONNECTION_STRING=postgres://postgres:postgres@localhost:5433/inventory_db?sslmode=disable
export TELEMETRY_DB_CONNECTION_STRING=postgres://postgres:postgres@localhost:5434/telemetry_db?sslmode=disable
```

### Installing `migrate`

Installation of the tool can be done using Go toolchain, the value of `-tags` is
used to install the correct database engine library with the tool.

> ```shell
>  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
> ```
>
> [Documentation][1]

## Available Commands

```shell
TARGETS:
  Makefile:db/start                      start database containers
  Makefile:migrations/up/shipment        run migrations for only shipment database
  Makefile:migrations/up/inventory       run migrations for only inventory database
  Makefile:migrations/up/telemetry       run migrations for only telemetry database
  Makefile:migrations/up                 run migrations for all databases
  Makefile:migrations/down/shipment      run teardown migrations for shipment database
  Makefile:migrations/down/inventory     run teardown migrations for inventory database
  Makefile:migrations/down/telemetry     run teardown migrations for telemetry database
  Makefile:migrations/down               run teardown migrations for all databases
  Makefile:targets                       print Makefile targets
```

### Usage

Run the following `make` commands as needed:

#### Start the Databases using docker-compose

```sh
make db/start
```

#### Stop the Databases using docker-compose

```sh
make db/stop
```

### Run Migrations (Up)

```sh
make migrations/up
```

Runs migrations to create database tables.

### Rollback Migrations (Down)

```sh
make migrations/down
```

Rolls back migrations, tearing down database tables.
