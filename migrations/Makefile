################################################################################
# Migrations Tasks
################################################################################
include ../.env
COMPOSE_FILE := ../manifests/docker-compose.yaml
DOCKER := podman

db/start: # @HELP start database containers in docker compose
	$(DOCKER) compose -f $(COMPOSE_FILE) up -d shipment-db telemetry-db

db/stop: # @HELP stop database containers in docker compose
	$(DOCKER) compose -f $(COMPOSE_FILE) down shipment-db telemetry-db

db/connect/shipment: # @HELP connect to shipment db
	psql ${SHIPMENT_DB_CONNECTION_STRING} 

db/connect/telemetry: # @HELP connect to telemetry db
	psql ${TELEMETRY_DB_CONNECTION_STRING} 

migrations/up/shipment: # @HELP run migrations for only shipment database
	migrate -path ./shipment -database ${SHIPMENT_DB_CONNECTION_STRING} up


migrations/up/telemetry: # @HELP run migrations for only telemetry database
	migrate -path ./telemetry -database ${TELEMETRY_DB_CONNECTION_STRING} up

migrations/up: # @HELP run migrations for all databases
migrations/up: migrations/up/shipment migrations/up/telemetry

migrations/down/shipment: # @HELP run teardown migrations for shipment database
	migrate -path ./shipment -database ${SHIPMENT_DB_CONNECTION_STRING} down 1

migrations/down/telemetry: # @HELP run teardown migrations for telemetry database
	migrate -path ./telemetry -database ${TELEMETRY_DB_CONNECTION_STRING} down 1

migrations/down: # @HELP run teardown migrations for all databases
migrations/down: migrations/down/telemetry migrations/down/shipment

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
