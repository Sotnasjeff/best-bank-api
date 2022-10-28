postgres:
	docker run --name postgres14.5 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.5-alpine

createdb:
	docker exec -it postgres14.5 createdb --username=root --owner=root best_bank

dropdb:
	docker exec -it postgres14.5 dropdb best_bank

migrateup:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5431/best_bank?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5431/best_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test