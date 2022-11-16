run:
	go run main.go

build:
	docker compose build --no-cache

up:
	docker compose up -d

down:
	docker compose down

db_setup:
	mysql -h 127.0.0.1 -u docker sampledb -p < ./repositories/testdata/setupDB.sql

db_in:
	mysql -h 127.0.0.1 -u docker sampledb -p

db_cleanup:
	mysql -h 127.0.0.1 -u docker sampledb -p < ./repositories/testdata/cleanupDB.sql

test: ## Execute tests
  ## go: -race requires cgo; enable cgo by setting CGO_ENABLED=1
	go test -race -shuffle=on ./...
