DB_URL=mysql://root:secret@tcp(localhost:3306)/invest

mysql:
	docker run --name mysql8 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=invest -d mysql:8

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

proto:
	rm -f proto/pb/*.go
	protoc --proto_path=proto/v1 --go_out=proto/pb --go_opt=paths=source_relative \
    --go-grpc_out=proto/pb --go-grpc_opt=paths=source_relative \
    proto/v1/*.proto

.PHONY: mysql new_migration migrateup migratedown proto