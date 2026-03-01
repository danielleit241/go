include .env
export

DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIRS) -seq $(NAME)

migrate-up:
	migrate -path $(MIGRATION_DIRS) -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path $(MIGRATION_DIRS) -database "$(DATABASE_URL)" down 1

migrate-down-n:
	migrate -path $(MIGRATION_DIRS) -database "$(DATABASE_URL)" down $(N)

migrate-force:
	migrate -path $(MIGRATION_DIRS) -database "$(DATABASE_URL)" force $(VERSION)

migrate-drop:
	migrate -path $(MIGRATION_DIRS) -database "$(DATABASE_URL)" drop

migrate-goto:
	migrate -path $(MIGRATION_DIRS) -database "$(DATABASE_URL)" goto $(VERSION)

.PHONY: migrate-create migrate-up migrate-down migrate-down-n