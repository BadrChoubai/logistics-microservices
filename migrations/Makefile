####################################################################################################
# Database Tasks
####################################################################################################

db/info: # @HELP prints value of DB_CONNECTION_STRING
db/info:
	@echo ${SHIPMENT_DB_CONNECTION_STRING}
	@echo ${INVENTORY_DB_CONNECTION_STRING}
	@echo ${TELEMETRY_DB_CONNECTION_STRING}


db/migrations/up/shipment:
	migrate -path ./shipment -database ${SHIPMENT_DB_CONNECTION_STRING} up

db/migrations/up/inventory:
	migrate -path ./inventory -database ${INVENTORY_DB_CONNECTION_STRING} up

db/migrations/up/telemetry:
	migrate -path ./telemetry -database ${TELEMETRY_DB_CONNECTION_STRING} up

db/migrations/up: db/migrations/up/shipment db/migrations/up/inventory db/migrations/up/telemetry

db/migrations/refresh: # @HELP refresh database tables
db/migrations/refresh: db/migrations/down db/migrations/up

