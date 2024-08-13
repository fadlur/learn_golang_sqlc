postgres:
	docker run --name postgres_local -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=147789 -d postgres:alpine

createdb:
	docker exec -it postgres_local createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres_local dropdb --username=postgres simple_bank

migration:
	migrate create -ext sql -dir db/migration --seq init_schema

migrateup: 
	migrate -path db/migration -database "postgresql://postgres:147789@localhost:5432/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:147789@localhost:5432/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
	
_PHONY: postgres createdb dropdb migrateup migratedown migration