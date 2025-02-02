# Base Stage with Go 1.23.5 (bookworm)
FROM golang:1.23.5-bookworm AS base
WORKDIR /app

# We install necessary dependencies (including SSL certificates)
RUN apt-get update && apt-get install -y --no-install-recommends \
    git ca-certificates libaom-dev \
    && rm -rf /var/lib/apt/lists/*

# Development stage
FROM base AS dev
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod \
    go install github.com/air-verse/air@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest

# We copy GO files for dependencies management and source code
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENTRYPOINT ["/bin/sh", "-c"]
CMD ["echo 'Container started. Use docker-compose to specify a command.'"]

# Compilation stage
FROM base AS builder
WORKDIR /build

# We copy the Go.Mod and Go.sum files for GO to manage the units
COPY --from=dev /app/go.mod /app/go.sum ./
# We copy the entire source code of the Development Stage Work Directory
COPY --from=dev /app ./

RUN CGO_ENABLED=0 go build -o go-events-guard

# Production stage with Distroless
FROM gcr.io/distroless/static-debian12 AS prod
WORKDIR /prod

# We copy the certificates by SSL
COPY --from=base /etc/ssl/certs /etc/ssl/certs

# We copy only the final executable
COPY --from=builder /build/go-events-guard ./
EXPOSE 8000
CMD ["/prod/go-events-guard"]
