# Invest API

## GOALs

- Set up structure
  - Docker
  - Kubernetes
  - AWS
  - Github
  - sh script
  - config
  - Makefile
- Design
  - DB
  - Apps
- Test
  - mockdb
  - CI
- API
  - gRPC
  - Auth
  - Async
  - Redis
  - Session Management
  - Logs
  - Error Handle
- Set up OpenAPI

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

## DEV STEPs

- Build a usable API
  - Apps Design
  - Structure set up (basic)
    - DB migration
    - Docker images
    - config
    - Makefile
  - API DEV (basic)
    - gRPC API
    - Error Handle
    - Auth
    - Logs
    - Async API
- Deploy
  - Docker fix
  - Kubernetes
  - AWS
    - RDS
    - secrets manager
    - EKS
    - Route53 (domain)
  - Auto deploy
- Better API
  - test (mockdb)
  - CI
  - OpenAPI
  - fix API
  - add more API

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
