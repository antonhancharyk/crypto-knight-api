# crypto-knight-api

API for the crypto-knight trading bot (tracks, entries, balance, orders, klines).

## Requirements

- Go 1.22+
- PostgreSQL
- gRPC common service (for enable/disable)
- Binance API keys (for balance, orders, position, klines)

## Environment variables

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_ENV` | Environment: `development` or `production` (affects logging) | `development` |
| `HTTP_ADDR` | HTTP server address | `:8080` |
| `DB_USER` | PostgreSQL user | — |
| `DB_PASSWORD` | PostgreSQL password | — |
| `DB_NAME` | PostgreSQL database name | — |
| `DB_HOST` | PostgreSQL host | — |
| `DB_PORT` | PostgreSQL port | — |
| `GRPC_HOST` | gRPC common service address | — |
| `PUBLIC_API_KEY` | Binance API key | — |
| `PRIVATE_API_KEY` | Binance API secret (signing) | — |

Copy `.env.example` to `.env` and fill in the values. RSA public key for auth is loaded from `./config/rsa/public_key.pem`.

## Run locally

```bash
cp .env.example .env
# edit .env

go run ./cmd
```

Server listens on `HTTP_ADDR` (default `:8080`).

## Tests

```bash
go test ./...
```

## API docs (Swagger)

With the server running, open:

- **Swagger UI:** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

Regenerate docs after changing handler annotations:

```bash
go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/main.go --parseDependency --parseInternal
```

## Endpoints (overview)

- `GET /healthz` — liveness
- `GET /common/status`, `GET /common/on`, `GET /common/off` — bot status and control
- `GET /auth/validate?token=...` — validate JWT
- `GET/POST /tracks`, `POST /tracks/bulk`, `GET /tracks/last` — tracks
- `GET/POST /tracks/history`, `POST /tracks/history/bulk` — tracks history
- `GET/POST /entries` — entries
- `GET /balance` — USDT balance
- `GET /orders` — open orders
- `GET /position` — positions
- `GET /klines?symbol=...` — klines

All routes except `/healthz` and `/swagger/*` require auth (query or header as configured).
