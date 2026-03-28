## Services

This section documents each service's data model to some depth.

### Shipment

The Shipment service owns the creation of shipments, status transitions, and route
information.

- Database: PostgreSQL – shipments are relational by nature, a shipment has a
  status history, waypoints, and references inventory items
- gRPC: Receives calls from Inventory when a shipment status changes to update
  stock levels; calls Inventory when a shipment is created to verify stock
  availability.

### Inventory

The inventory service owns the item catalog, stock levels, and warehouse
locations. 

- Database: PostgreSQL – inventory is inherently relational, items have
  categories, locations, and stock thresholds
- gRPC: calls Shipment to update inventory when a shipment status changes;
  responds to Shipment when stock availability is checked

#### Data Model

```mermaid
erDiagram
    SHIPMENT {
        uuid        id PK
        timestamp   order_date
        VARCHAR(50) status
        VARCHAR(50) origin
        VARCHAR(50) destination
        timestamp   created_at
        timestamp   updated_at
    }

    CONTAINER {
        uuid        id PK
        uuid        shipment_id FK
        VARCHAR(11) container_number
        timestamp   created_at
        timestamp   updated_at

    }

    SHIPMENT_ITEM {
        uuid        id PK
        uuid        shipment_id FK
        uuid        inventory_id FK
        int         quantity
        timestamp   created_at
        timestamp   updated_at
    }

    INVENTORY {
        uuid        id PK
        uuid        warehouse_id FK
        VARCHAR(50) name
        int         quantity
        timestamp   created_at
        timestamp   updated_at
    }

    WAREHOUSE {
        uuid        id PK
        VARCHAR(50) address
        VARCHAR(50) city
        VARCHAR(50) state
        VARCHAR(50) zip_code
        timestamp   created_at
        timestamp   updated_at
    }

    SHIPMENT ||--o{ CONTAINER : "has"
    SHIPMENT ||--o{ SHIPMENT_ITEM : "contains"
    SHIPMENT_ITEM }o--|| INVENTORY : "references"
    INVENTORY }o--|| WAREHOUSE : "stored in"

```


### Telemetry

The Telemetry service owns sensor reading data attached to shipments:
temperature, humidity, GPS coordinates.

- Database: PostgreSQL – chosen for operational simplicity; a production system
  might use TimescaleDB or InfluxDB for time-series query performance at scale
- No gRPC—telemetry is write-heavy and read-only from other services'
  perspective, no service-to-service calls needed (in V1).

#### Data Model

```mermaid
erDiagram
    SENSOR {
        uuid            id PK
        uuid            container_id FK
        ENUM            sensor_type "(Environmental, Location, Security)"     
        timestamp       created_at
        timestamp       updated_at

    }

    ENVIRONMENTAL_READING {
        uuid            id PK
        uuid            sensor_id FK
        int             temperature_celsius
        timestamp       recorded_at
        timestamp       created_at
    }

    LOCATION_READING {
        uuid            id PK
        uuid            sensor_id FK
        DECIMAL         latitude "`latitude` is `DECIMAL(9,6)` in the database schema."
        DECIMAL         longitude "`longitude` is `DECIMAL(9,6)` in the database schema."
        timestamp       recorded_at
        timestamp       created_at
    }

    SECURITY_READING {
        uuid            id PK
        uuid            sensor_id FK
        boolean         door_open
        timestamp       recorded_at
        timestamp       created_at
    }

    SENSOR ||--o{ ENVIRONMENTAL_READING : "produces"
    SENSOR ||--o{ LOCATION_READING : "produces"
    SENSOR ||--o{ SECURITY_READING : "produces"
```

