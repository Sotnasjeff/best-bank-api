DB_URL=postgresql://root:secret@localhost:5432/best_bank?sslmode=disable
network:
	docker network create -d bridge bank-network

container:
	docker run --name best_bank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres14.5:5432/best_bank?sslmode=disable" best_bank:latest

postgres:
	docker run --name postgres14.5 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.5-alpine

createdb:
	docker exec -it postgres14.5 createdb --username=root --owner=root best_bank

dropdb:
	docker exec -it postgres14.5 dropdb best_bank

migrateawsup:
	migrate --path db/migration --database "postgresql://root:F8hJAOmYIcpiPp0zUdz0@best-bank.ctmiyuptewzy.us-east-1.rds.amazonaws.com:5432/best_bank" --verbose up

migrateup:
	migrate --path db/migration --database "$(DB_URL)" --verbose up

migrateup1:
	migrate --path db/migration --database "$(DB_URL)" --verbose up 1

migratedown:
	migrate --path db/migration --database "$(DB_URL)" --verbose down

migratedown1:
	migrate --path db/migration --database "$(DB_URL)" --verbose down 1

db_docs:
	sudo dbdocs build doc/db.dbml

db_schema:
	sudo sudo dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server: 
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/best-bank-api/db/sqlc Store

createmigration:
	migrate create -ext sql -dir db/migration -seq add_users

proto: 
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=best_bank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: createdb dropdb postgres migrateup migratedown migrateup1 migratedown1 sqlc test server mock createmigration network container migrateawsup db_docs db_schema proto evans redis