run: ## TBU
	go run main.go

build: ## TBU
	docker compose build

up: ## TBU
	docker compose up -d

up_only_app: ## Confirm error logs in stout when db does not run
	docker compose up app

down: ## TBU
	docker compose down

db_setup: ## TBU
	mysql -h 127.0.0.1 -u docker sampledb -p < ./repositories/testdata/setupDB.sql

db_in: ## TBU
	mysql -h 127.0.0.1 -u docker sampledb -p

db_cleanup: ## TBU
	mysql -h 127.0.0.1 -u docker sampledb -p < ./repositories/testdata/cleanupDB.sql

test: ## Execute tests
  ## go: -race requires cgo; enable cgo by setting CGO_ENABLED=1
	go test -race -shuffle=on ./...

google_pkey: ## show google's pulic keys
	http https://accounts.google.com/.well-known/openid-configuration

google_jwt_pkey: ## show google's pulic key for issuing jwt
	http https://www.googleapis.com/oauth2/v3/certs

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
