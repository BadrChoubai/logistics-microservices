# Logistics Microservices

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
│   ├── inventory
│   ├── shipment
│   └── telemetry
├── internal
│   └── gateway
├── manifests
│   ├── Dockerfile
│   ├── docker-compose.yaml
│   ├── gateway
│   ├── inventory
│   ├── shipment
│   └── telemetry
└── migrations
    ├── Makefile
    ├── README.md
    ├── inventory
    ├── shipment
    └── telemetry
```

---

## 🚀 Running the Project

### Prerequisites

- Go 1.25+
- Docker & Docker Compose
- Make (recommended)


### 🔐 Environment Variables

Environment variables are managed via `.env`

```bash
SHIPMENT_DB_CONNECTION_STRING=postgres://postgres:postgres@localhost:5432/shipment_db?sslmode=disable
INVENTORY_DB_CONNECTION_STRING=postgres://postgres:postgres@localhost:5433/inventory_db?sslmode=disable
TELEMETRY_DB_CONNECTION_STRING=postgres://postgres:postgres@localhost:5434/telemetry_db?sslmode=disable
```

> ⚠️ Do not commit `.env` files (they are ignored by `.gitignore`)

---

### 1. Build & Run with Docker Compose

```bash
docker compose -f manifests/docker-compose.yaml up --build
```

Or run in detached mode:

```bash
docker compose -f manifests/docker-compose.yaml up -d --build
```

---

### 2. Stop Services

```bash
docker compose -f manifests/docker-compose.yaml down
```

---

### 3. View Logs

```bash
docker compose -f manifests/docker-compose.yaml logs -f gateway
```

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
