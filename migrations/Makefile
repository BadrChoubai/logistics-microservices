################################################################################
# Migrations Tasks
################################################################################
include ../.env

db/start: # @HELP start database containers in docker compose
	docker compose -f ../manifests/docker-compose.yaml up -d logistics-db telemetry-db

db/stop: # @HELP stop database containers in docker compose
	docker compose -f ../manifests/docker-compose.yaml down logistics-db telemetry-db

migrations/up/logistics: # @HELP run migrations for only logistics database
	migrate -path ./logistics -database ${LOGISTICS_DB_CONNECTION_STRING} up


migrations/up/telemetry: # @HELP run migrations for only telemetry database
	migrate -path ./telemetry -database ${TELEMETRY_DB_CONNECTION_STRING} up

migrations/up: # @HELP run migrations for all databases
migrations/up: migrations/up/logistics migrations/up/telemetry

migrations/down/logistics: # @HELP run teardown migrations for logistics database
	migrate -path ./logistics -database ${LOGISTICS_DB_CONNECTION_STRING} down 1

migrations/down/telemetry: # @HELP run teardown migrations for telemetry database
	migrate -path ./telemetry -database ${TELEMETRY_DB_CONNECTION_STRING} down 1

migrations/down: # @HELP run teardown migrations for all databases
migrations/down: migrations/down/logistics migrations/down/telemetry

targets: # @HELP print Makefile targets
	echo "TARGETS:"
	grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)     \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s	%s\n", $$1, $$2 };  \
	    '

help: targets

.DEFAULT_GOAL=help
.SILENT: help targets
.PHONY: db/start db/stop targets migrations/down migrations/up migrations/up/% migrations/down/%
