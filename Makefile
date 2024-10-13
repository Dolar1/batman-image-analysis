# Makefile for the Batman project

.PHONY: run

# Command to run the Go application
run:
	go run cmd/main.go start

db-migration:
	go run cmd/main.go migrate

revert-db-migration:
	go run cmd/main.go migrate:down

runair:
	air