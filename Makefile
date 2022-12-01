postgres:
	docker run --name postgres14.5 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.5-alpine

createdb:
	docker exec -it postgres14.5 createdb --username=root --owner=root best_bank

dropdb:
	docker exec -it postgres14.5 dropdb best_bank

migrateup:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/best_bank?sslmode=disable" --verbose up

migrateup1:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/best_bank?sslmode=disable" --verbose up 1

migratedown:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/best_bank?sslmode=disable" --verbose down

migratedown1:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/best_bank?sslmode=disable" --verbose down 1

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

.PHONY: createdb dropdb postgres migrateup migratedown migrateup1 migratedown1 sqlc test server mock createmigration