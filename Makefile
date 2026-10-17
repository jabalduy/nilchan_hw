include .env
export

.PHONY: service-run
service-run:
	@go run .

.PHONY: migrate-up
migrate-up:
	@migrate -path migrations \
	-database ${CONN_STRING} up

.PHONY: migrate-down
migrate-down:
	@migrate -path migrations \
	-database ${CONN_STRING} down