# Fitness App - Clean Architecture (Go)

This template contains:
- Clean Architecture layout (domain, usecase, infrastructure, interface)
- GORM + Postgres setup
- CRUD REST endpoints for User, WorkoutSet, Message
- Docker + docker-compose for local Postgres
- WhatsApp webhook handler placeholder (for WhatsApp Cloud API or Twilio)
- Mermaid diagram describing the architecture

## How to use

1. Copy `.env.example` to `.env` and edit DB credentials.
2. Start Postgres with Docker Compose:
   ```
   docker-compose up -d
   ```
3. Build & run locally:
   ```
   go build ./cmd/api
   ./api
   ```
   OR build inside Docker (see Dockerfile).
4. Expose webhook for WhatsApp using ngrok or similar and configure the provider.

## Notes
- Replace `github.com/yourusername/fitness-app` in `go.mod` with your module path.
- The WhatsApp handler is a placeholder: configure it with your provider (Facebook/Meta WhatsApp Cloud API or Twilio).
