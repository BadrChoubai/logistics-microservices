FROM golang:1.25 AS builder

ARG OS="linux"
ARG ARCH="amd64"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH}

RUN go build -o /bin/inventory ./cmd/inventory

FROM gcr.io/distroless/base-debian12

COPY --from=builder /bin/inventory /app
COPY --from=builder /app/manifests/inventory/config.json /config.json
EXPOSE 8080

ENTRYPOINT ["/app"]
