# Ragnarok-multifinance

A high-performance, PCI DSS-compliant Golang microservice for unified installment management, built with Gin and clean hexagonal architecture.

## Features
- Product management with flexible tenures
- Transaction processing with Redis session validation
- Automatic payment schedule generation
- Secure JWT authentication and input validation
- Proactive error handling and monitoring
- Graceful shutdown and zero-downtime deployments

## Architecture
```
ragnarok-multifinance/
├── internal/
│   ├── model/
│   │   ├── request/
│   │   ├── database/
│   │   └── response/
│   ├── usecase/
│   ├── delivery/
│   │   └── http/
│   ├── repository/
│   │   └── userRepo/
│   ├── library/
│   │   ├── mysql/
│   │   └── redis/
│   └── pkg/
├── config/
│   └── config.go
├── Dockerfile
├── go.mod
├── main.go
```

## Quickstart

1. **Clone the repo**
2. **Configure MySQL and Redis** (see `config/config.json`)
3. **Build and run**
   ```sh
   go mod tidy
   go build -o ragnarok .
   ./ragnarok
   ```
4. **Docker**
   ```sh
   docker build -t ragnarok-multifinance .
   docker run -p 8080:8080 ragnarok-multifinance
   ```

## Endpoints
- `GET /products` - List products
- `POST /products` - Create product (admin)
- `POST /transactions` - Create transaction
- `POST /payments` - Record payment
- `GET /transactions/{id}/schedule` - Payment schedule
- `GET /health` - Health check
- `GET /ready` - Readiness check

## Security
- JWT authentication
- Redis session validation
- Input validation (go-playground/validator)
- Secure headers (gin-contrib/secure)

## License
MIT 