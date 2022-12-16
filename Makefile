migrateup:
	migrate -path db/migration -database "postgresql://@85.214.139.19:7777/bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://@85.214.139.19:7777/bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://@85.214.139.19:7777/bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://@85.214.139.19:7777/bank?sslmode=disable" -verbose down 1

migratecreate:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: migratecreate migrateup migratedown migrateup1 migratedown1 sqlc test  server