PG_DSN?="postgres://localhost:5432/database?sslmode=disable"

db-help:
	@echo '  ${PG_DSN}'
	@echo '  '
	@echo '  make db-status'
	@echo '  make db-up'
	@echo '  make db-down'
	@echo '  make db-seed'
	@echo ' '
	@echo '  make NAME="create_pages" db-create'
	@echo '  make db-upgrade'
	@echo '  '

db-status:
	goose -dir migrations postgres "$(PG_DSN)" status

db-create: NAME=$NAME
db-create:
	goose -dir migrations postgres "$(PG_DSN)" create $(NAME) sql

db-up:
	goose -dir migrations postgres "$(PG_DSN)" up

db-down:
	goose -dir migrations postgres "$(PG_DSN)" down

db-seed:
	goose -dir seeds -table goose_seed_version postgres "$(PG_DSN)" up

db-upgrade:
	curl https://raw.githubusercontent.com/shlima/oi/master/db.mk -O
