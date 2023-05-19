include config/app.env
DB_URL="mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}"

server:
	go run main.go

mysql:
	docker run --name mysql -p ${DB_PORT}:${DB_PORT} -e MYSQL_ROOT_PASSWORD=${DB_PASSWORD} -e MYSQL_DATABASE=${DB_NAME} -d mysql:8

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database $(DB_URL) -verbose up

migratedown:
	migrate -path db/migration -database $(DB_URL) -verbose down

gen_pb:
	rm -f proto/pb/*.go
	protoc --proto_path=proto/v1 --go_out=proto/pb --go_opt=paths=source_relative \
    --go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative \
    proto/v1/*.proto

.PHONY: server mysql new_migration migrateup migratedown gen_pb