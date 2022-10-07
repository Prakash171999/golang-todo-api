include .env
MIGRATE=docker-compose exec web migrate -path=migration -database "mysql://${DBUsername}:${DBPassword}@tcp(${DBHost}:${DBPort})/${DBName}" -verbose

dev:
		gin appPort ${ServerPort} -i run server.go
migrate-up:
		$(MIGRATE) up
migrate-down:
		$(MIGRATE) down 
force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

drop:
		$(MIGRATE) drop

create:
		@read -p  "What is the name of migration?" NAME; \
		${MIGRATE} create -ext sql -seq -dir migration  $$NAME

crud:
	bash automate/scripts/crud.sh

.PHONY: migrate-up migrate-down force goto drop create

.PHONY: migrate-up migrate-down force goto drop create auto-create


# Migration commands for local server not in docker container
MIGRATE_LOCAL=migrate -source file://migration -database "mysql://${DBUsername}:${DBPassword}@tcp(${DBHost}:${DBPort})/${DBName}" -verbose

migrate-up-local:
		$(MIGRATE_LOCAL) up
migrate-down-local:
		$(MIGRATE_LOCAL) down 
force-local:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE_LOCAL) force $$VERSION

goto-local:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE_LOCAL) goto $$VERSION

drop-local:
		$(MIGRATE_LOCAL) drop

create-local:
		@read -p  "What is the name of migration?" NAME; \
		${MIGRATE_LOCAL} create -ext sql -seq -dir migration  $$NAME

crud-local:
	bash automate/scripts/crud.sh