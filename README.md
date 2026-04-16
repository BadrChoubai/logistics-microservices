# _Logistique_

This repository contains a Go-based microservices architecture with an API Gateway as the central entry point. The project uses Docker and Docker Compose for local development and service orchestration.

---

## 🧱 Project Structure

```text
.
├── .env
├── LICENSE
├── Makefile
├── README.md
├── api
│   └── swagger
├── cmd
│   ├── gateway
│   ├── shipment
│   └── telemetry
├── internal
│   └── gateway
├── manifests
│   ├── docker-compose.yaml
│   ├── gateway
│   │   └── gateway
│   ├── shipment
│   │   └── Dockerfile
│   └── telemetry
│       └── Dockerfile
└── migrations
    ├── Makefile
    ├── logistics
    └── telemetry
```

---

## 🚀 Running the Project

### Prerequisites

- Go 1.25+
- Docker & Docker Compose
- Make (recommended)

### Development Tools

#### `golangci-lint`

Used to run linters

```bash
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.11.4
```

#### `swag`

Used to generate API docs

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

#### `migrate`

Used to run database migrations

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 🔐 Environment Variables

Environment variables are managed via `.env`

```bash
LOGISTICS_DB_CONNECTION_STRING=postgres://postgres:postgres@localhost:5432/logistics_db?sslmode=disable
TELEMETRY_DB_CONNECTION_STRING=postgres://postgres:postgres@localhost:5434/telemetry_db?sslmode=disable
```

> ⚠️ Do not commit `.env` files (they are ignored by `.gitignore`)

---

### 1. Build & Run with Docker Compose

```bash
[docker|podman] compose -f manifests/docker-compose.yaml up --build
```

Or run in detached mode:

```bash
[docker|podman] compose -f manifests/docker-compose.yaml up -d --build
```

---

### 2. Stop Services

```bash
[docker|podman] compose -f manifests/docker-compose.yaml down
```

---

### 3. View Logs

```bash
[docker | podman] compose -f manifests/docker-compose.yaml logs -f gateway
```

## 🗄️ Database Access

Connection strings are defined in `.env` (see above). For local development with a GUI client:

| Database     | URL                                             |
| ------------ | ----------------------------------------------- |
| logistics_db | `jdbc:postgresql://localhost:5432/logistics_db` |
| telemetry_db | `jdbc:postgresql://localhost:5434/telemetry_db` |

**User:** `postgres` / **Password:** `postgres`

> SSL is disabled for local development (`sslmode=disable`).

### DataGrip

1. **+** → **Data Source from URL** → paste the URL from the table above
2. Enter credentials, click **Download missing driver files** if prompted
3. **Test Connection** → **OK**

---

## 📦 Docker Image Versioning

This project uses Git-based versioning for Docker images.

- Each image is tagged using the current Git commit:

  ```
  logistics-gateway:<git-sha>
  ```

- You can inspect the version with:

```bash
make version
```

Example output:

```
b3d2787
```

---

## 🧪 Development Workflow

1. Make changes to the code
2. Build a new versioned image:

   ```bash
   make image
   ```

3. Run services:

   ```bash
   docker compose -f manifests/docker-compose.yaml up --build
   ```

_Other Commands_:

```bash
make build        # compile all binaries to ./bin
make ci           # run lint + tests
make docs         # regenerate Swagger spec after changing annotations
```

---

## 📖 API Documentation

Swagger documentation is the **single source of truth** for all APIs.

Generated docs are located in:

```
api/swagger/
```

Access the API documentation at:

```
http://localhost:8080/api
```

---

## 🧰 Makefile Commands

Common commands:

```bash
make build        # Build binaries locally
make image        # Build Docker images
make clean        # Clean build artifacts and images
make test         # Run tests
make lint         # Run linters
make version      # Print current version
```

---

## 🧠 Notes

- The API Gateway acts as the **entry point and source of truth** for API documentation.
- Services are containerized and orchestrated via Docker Compose.
- Git commit hashes are used for deterministic versioning.
- The project is structured as a **monorepo** containing multiple microservices.

---

## 📌 Future Improvements

[//]: # "TODO: Create Project Board"

---

## 📜 License

This project is licensed under the MIT License.
