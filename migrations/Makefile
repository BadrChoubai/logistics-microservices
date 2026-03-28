################################################################################
# Migrations Tasks
################################################################################
include ../.env

db/start: # @HELP start database containers
	docker compose -f ../manifests/docker-compose.yaml up -d shipment-db inventory-db telemetry-db

db/stop:
	docker compose -f ../manifests/docker-compose.yaml down shipment-db inventory-db telemetry-db

migrations/up/shipment: # @HELP run migrations for only shipment database
	migrate -path ./shipment -database ${SHIPMENT_DB_CONNECTION_STRING} up

migrations/up/inventory: # @HELP run migrations for only inventory database
	migrate -path ./inventory -database ${INVENTORY_DB_CONNECTION_STRING} up

migrations/up/telemetry: # @HELP run migrations for only telemetry database
	migrate -path ./telemetry -database ${TELEMETRY_DB_CONNECTION_STRING} up

migrations/up: # @HELP run migrations for all databases
migrations/up: migrations/up/shipment migrations/up/inventory migrations/up/telemetry

migrations/down/shipment: # @HELP run teardown migrations for shipment database
	migrate -path ./shipment -database ${SHIPMENT_DB_CONNECTION_STRING} down 1

migrations/down/inventory: # @HELP run teardown migrations for inventory database
	migrate -path ./inventory -database ${INVENTORY_DB_CONNECTION_STRING} down 1

migrations/down/telemetry: # @HELP run teardown migrations for telemetry database
	migrate -path ./telemetry -database ${TELEMETRY_DB_CONNECTION_STRING} down 1

migrations/down: # @HELP run teardown migrations for all databases
migrations/down: migrations/down/shipment migrations/down/inventory migrations/down/telemetry

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
.PHONY: db/start db/stop targets migrations/down migrations/up migrations/up/%
