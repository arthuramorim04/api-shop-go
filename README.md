# Go Shop API

API em Go que replica o comportamento do projeto Node/Express existente em `shop-api/`.

## Stack
- Gin (HTTP)
- MySQL (go-sql-driver)
- JWT (golang-jwt)
- Bcrypt (x/crypto)
- Mercado Pago (REST API)

## Endpoints (equivalentes)
- Auth
  - POST /login
- Users (proteção: auth + role support)
  - GET /users
  - POST /users
  - GET /users/:id
  - PUT /users/:id
  - DELETE /users/:id
- Products
  - POST /products (auth + admin)
  - GET /products (público)
  - GET /products/:productId (auth + admin)
  - PUT /products/:productId (auth + admin)
  - DELETE /products/:productId (auth + admin)
- Orders
  - POST /orders (auth + admin)
  - GET /orders (auth + support)
  - GET /orders/:id (auth + support)
  - GET /orders/user/:userId (auth)
  - GET /orders/product/:productId (auth + support)
  - PATCH /orders/:id/status (auth + admin)
- Payment
  - POST /payment/notification (webhook Mercado Pago)

## Variáveis de ambiente (.env)
```
APP_PORT=8081
DB_DSN=user:password@tcp(127.0.0.1:3306)/shopdb?parseTime=true&charset=utf8mb4
JWT_SECRET=supersecret
MP_ACCESS_TOKEN=seu_token_mp
MP_NOTIFICATION_URL=http://localhost:8081
```

## Rodando
```
go mod tidy
go run ./...
```

## Linting
Este projeto usa `golangci-lint` com a configuração em `.golangci.yml` na raiz.

### Instalação (Windows)
- Scoop:
  ```
  scoop install golangci-lint
  ```
- Chocolatey:
  ```
  choco install golangci-lint
  ```
- Via Go (binário no GOPATH/bin):
  ```
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

### Rodando o linter
```
golangci-lint run ./...
```

### Correções automáticas (quando disponíveis)
```
golangci-lint run --fix
```

