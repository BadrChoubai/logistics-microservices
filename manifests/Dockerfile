FROM golang:1.25 AS builder

ARG SERVICE
ARG OS="linux"
ARG ARCH="amd64"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=${OS} GOARCH=${ARCH}

RUN go build -o /bin/${SERVICE} ./cmd/${SERVICE}

FROM gcr.io/distroless/base-debian12
ARG SERVICE

COPY --from=builder /bin/${SERVICE} /app
EXPOSE 8080

ENTRYPOINT ["/app"]
