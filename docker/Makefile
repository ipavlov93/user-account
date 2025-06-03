include ./app/.env

BINARY_NAME=app

# Makefile has list of following migrate commands using goose cli:

GOOSE_ENV=GOOSE_MIGRATION_DIR=${GOOSE_MIGRATION_DIR} GOOSE_DRIVER=${GOOSE_DRIVER} GOOSE_DBSTRING=${GOOSE_DBSTRING}

goose_install:
	go install github.com/pressly/goose/v3@latest

# always create new migration file and optionally apply sequential ordering to it (if previous files exist)"
create_valid_migration: create_migration migrate_fix migrate_validate

create_migration:
	@echo "Running goose create init sql"
	$(GOOSE_ENV) goose create new_migration sql

migrate_version:
	$(GOOSE_ENV) goose version

migrate_status:
	$(GOOSE_ENV) goose status

migrate_fix:
	@echo "Running goose fix to apply sequential ordering to migrations"
	$(GOOSE_ENV) goose fix

# migrate up with timeout
migrate_up:
	@echo "Running goose up to migrate the DB to the most recent version available"
	$(GOOSE_ENV) goose up

# migrate up-by-one with timeout
migrate_up_by_one:
	@echo "Running goose up to migrate the DB up-by-one"
	$(GOOSE_ENV) goose up-by-one

# migrate down-by-one with timeout
migrate_down_by_one:
	@echo "Running goose down to rollback the DB down-by-one"
	$(GOOSE_ENV) goose down

migrate_validate:
	@echo "Running goose validate migration files (goose syntax only) without running them"
	$(GOOSE_ENV) goose validate
