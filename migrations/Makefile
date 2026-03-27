################################################################################
# Migrations Tasks
################################################################################

# Do not set these here, they should be in .env at root of project
SHIPMENT_DB_CONNECTION_STRING ?= 
INVENTORY_DB_CONNECTION_STRING ?=
TELEMETRY_DB_CONNECTION_STRING ?=

info: # @HELP prints all database connections 
	@echo "VARIABLES:"
	@printf "  %-30s  %s\n" SHIPMENT_DB_CONNECTION_STRING $(SHIPMENT_DB_CONNECTION_STRING)
	@printf "  %-30s  %s\n" INVENTORY_DB_CONNECTION_STRING $(INVENTORY_DB_CONNECTION_STRING) 
	@printf "  %-30s  %s\n" TELEMETRY_DB_CONNECTION_STRING $(TELEMETRY_DB_CONNECTION_STRING)


migrations/up/shipment: # @HELP run migrations for only shipment database
	migrate -path ./shipment -database ${SHIPMENT_DB_CONNECTION_STRING} up

migrations/up/inventory: # @HELP run migrations for only inventory database
	migrate -path ./inventory -database ${INVENTORY_DB_CONNECTION_STRING} up

migrations/up/telemetry: # @HELP run migrations for only telemetry database
	migrate -path ./telemetry -database ${TELEMETRY_DB_CONNECTION_STRING} up

migrations/up: # @HELP run migrations for all databases
migrations/up: migrations/up/shipment  migrations/up/inventory migrations/up/telemetry

migrations/down/shipment: # @HELP run teardown migrations for only shipment database
	migrate -path ./shipment -database ${SHIPMENT_DB_CONNECTION_STRING} 

migrations/down/inventory: # @HELP run teardown migrations for only inventory database
	migrate -path ./inventory -database ${INVENTORY_DB_CONNECTION_STRING} down

migrations/down/telemetry: # @HELP run teardown migrations for only telemetry database
	migrate -path ./telemetry -database ${TELEMETRY_DB_CONNECTION_STRING} down

migrations/down: # @HELP run teardown migrations for all databases
migrations/down: migrations/down/shipment migrations/down/inventory migrations/down/telemetry

targets: # @HELP print Makefile targets
	echo "TARGETS:"
	grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST)     \
	    | awk '                                   \
	        BEGIN {FS = ": *# *@HELP"};           \
	        { printf "  %-30s %s\n", $$1, $$2 };  \
	    '

help: info targets

.DEFAULT_GOAL=help
.SILENT: help info targets
.PHONY: all info targets migrations/down migrations/down/% migrations/up migrations/up/%
