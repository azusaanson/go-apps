# Invest gRPC API

## Get Started

1. Install golang-migrate

```bash
brew install golang-migrate
```

2. Build mysql container and run migrate up

```bash
make mysql
```

```bash
make migrateup
```

3. Start gRPC server

```bash
make server
```

## TECH

- Docker
- Kubernetes
- AWS
  - RDS
  - secrets manager
  - EKS
  - Route53 (domain)
- OpenAPI
- CircleCI
- MySQL
- Gorm
- GIN
- testing
- gomock
- gRPC
- Redis
- PASETO
