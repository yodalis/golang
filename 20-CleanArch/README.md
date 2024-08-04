# Clean Architecture Challenge

Welcome to my Clean Architecture Challenge!

## Table of Contents

- [Ports](#ports)
- [Commands](#commands)

## Ports

Here is the list of ports where GRPC, GraphQL, and the Web Server are allocated.

- GRPC: 50051
- GraphQL: 8080
- Web Server: 8000

## Commands

Steps to run the project locally.

```bash
# Init docker
docker-compose up -d
```

```bash
# Enter MySQL bash
mysql -u root -p
```

```bash
# Run migrations
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
```

```bash
# Destroy all from database (migration)
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down
```

```bash
# Run application:
# (in case you are in the root package of the project)
go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

# in case you are inside cmd/ordersystem
go run main.go wire_gen.go

```

```bash
# Run evans to GRPC
evans --proto internal/infra/grpc/protofiles/order.proto repl
```

#### GraphQL can be accessed by running the server and entering localhost:8080 in the browser

