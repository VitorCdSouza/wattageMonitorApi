postgres:
	docker run --name postgresql16 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:16-alpine

createdb:
	docker exec -it postgresql16 createdb --username=root --owner=root wattage_monitor

dropdb:
	docker exec -it postgresql16 dropdb wattage_monitor

migrateup:
	migrate -path db/migration -database "postgresql://root:1234@localhost:5433/wattage_monitor?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:1234@localhost:5433/wattage_monitor?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres dropdb createdb migrateup migratedown sqlc
