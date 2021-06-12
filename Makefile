build-migration:
	go build -o migration cmd/migrate/main.go

migration-create: build-migration
	./migration create $(NAME)

migration-up: build-migration
	./migration up

migration-down: build-migration
	./migration down

migration-redo: build-migration
	./migration redo

migration-status: build-migration
	./migration status

migration-version: build-migration
	./migration version

migration-fix: build-migration
	./migration fix

build:
	go build -o api cmd/api/main.go

run: build
	./api

.PHONY: build run