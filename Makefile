MIGRATIONS_PATH = internal/infrastructure/db/migrations
DSN = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

include config/.env.$(ENV)

app.run:
	go run cmd/app/main.go

migration.new:
	migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(name)

migration.up:
	migrate -verbose -database $(DSN) -path $(MIGRATIONS_PATH) up

migration.goto:
	migrate -database $(DSN) -path $(MIGRATIONS_PATH) goto $(version)

migration.version:
	migrate -database $(DSN) -path $(MIGRATIONS_PATH) version