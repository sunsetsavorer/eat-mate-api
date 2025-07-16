MIGRATIONS_PATH = internal/infrastructure/db/migrations

include config/.env

app.run:
	go run cmd/app/main.go

migration.new:
	migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(name)

migration.up:
	migrate -verbose -database $(DBCONN) -path $(MIGRATIONS_PATH) up

migration.goto:
	migrate -database $(DBCONN) -path $(MIGRATIONS_PATH) goto $(version)

migration.version:
	migrate -database $(DBCONN) -path $(MIGRATIONS_PATH) version