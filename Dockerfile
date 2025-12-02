FROM golang:1.20-alpine AS build
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api

FROM alpine:latest
WORKDIR /
COPY --from=build /api /api
COPY .env.example /app/.env.example
EXPOSE 8080
CMD ["/api"]
