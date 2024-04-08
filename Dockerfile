## Build stage
FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


## Deploy stage
FROM alpine:latest

WORKDIR /root/

# for SSL connection
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

CMD ["./main"]
