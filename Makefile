current_dir := $(shell pwd)

.PHONY: gen-docs
gen-docs:
	docker run --rm -v "$(current_dir):/code" \
		ghcr.io/swaggo/swag:v1.16.4 init \
		-g ./api/main.go \
		-d cmd,internal \
		-o internal/docs
	docker run --rm -v "$(current_dir):/code" \
		ghcr.io/swaggo/swag:v1.16.4 fmt \
		-d internal/docs
	docker run --rm -v "$(current_dir):/code" \
		ghcr.io/swaggo/swag:v1.16.4 fmt \
		-d .

.PHONY: reset-all
reset-all:
	docker compose down -v
	docker compose up -d

.PHONY: regenerate-db
regenerate-db:
	docker compose down -v db
	docker compose up -d db

.PHONY: seed
seed:
	docker compose exec -T db bash -c "until pg_isready -U $${POSTGRES_USER:-postgres} -d $${POSTGRES_DB:-band-organizer}; do sleep 1; done"
	@go run cmd/seed/main.go