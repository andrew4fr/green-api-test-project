# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /build

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /green-api-test-project .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /green-api-test-project .

ENV GREEN_API_URL=https://3100.api.green-api.com
ENV SERVER_ADDR=:8080

EXPOSE 8080

CMD ["/app/green-api-test-project"]
