## Build stage
FROM golang:1.24.1-bullseye as base

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download

## Development stage
FROM base AS development

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/air-verse/air@latest

COPY . .

WORKDIR /app/cmd
RUN mkdir -p /app/tmp

CMD ["air", "-c", ".air.toml"]
