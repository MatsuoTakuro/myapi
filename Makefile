run:
	go run main.go

db_up:
	docker compose up -d

db_down:
	docker compose down

db_init:
	mysql -h 127.0.0.1 -u docker sampledb -p < createTable.sql

db_in:
	mysql -h 127.0.0.1 -u docker sampledb -p

test: ## Execute tests
  ## go: -race requires cgo; enable cgo by setting CGO_ENABLED=1
	go test -race -shuffle=on ./...
