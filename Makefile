all: start

start: migration_up run

run:
	go run app/app.go

migration_up: 
	migrate -path migrations/ -database="postgresql://elmir:user@localhost:5433/shortlink?sslmode=disable" up

migration_down: 
	migrate -path migrations/ -database="postgresql://elmir:user@localhost:5433/shortlink?sslmode=disable" down

lint:
	golangci-lint run 

format:
	find. -name "*.go" -exec go fmt {} \;

.PHONY: lint format migration_up migration_down run start
