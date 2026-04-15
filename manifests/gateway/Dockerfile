FROM golang:1.25 AS builder

ARG OS="linux"
ARG ARCH="amd64"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH}

RUN go build -o /bin/gateway ./cmd/gateway

FROM gcr.io/distroless/base-debian12

COPY --from=builder /bin/gateway /app/gateway
COPY --from=builder /app/manifests/gateway/config.json /config.json
EXPOSE 8080

ENTRYPOINT ["/app/gateway"]
