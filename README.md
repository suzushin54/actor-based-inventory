# actor-based-inventory

This is a simple inventory system that uses actors to manage the state of the inventory. 

The inventory is implemented using the actor model, which allows for concurrent access to the inventory state.

## Technologies
- [Go](https://golang.org/)
- [Proto.Actor](https://proto.actor/) - A cross-platform actors framework
- [Kafka](https://kafka.apache.org/) - Used for building real-time data pipelines and streaming apps
- [Cassandra](http://cassandra.apache.org/) - A distributed NoSQL database
- [Buf Connect](https://docs.buf.build/connect/introduction) - Used for generating and managing gRPC and HTTP API layers from protobuf definitions

## Development

- `make setup`: Install necessary tools (Buf, grpcurl, connect-go, Wire) and tidy up Go modules.
- `make install-check`: Check if all necessary tools are correctly installed.
- `make buf`: Lint and generate code from protobuf files using Buf.
- `make wire`: Generate dependency injection code using Google Wire.
- `make test`: Run the scenario test using the `runn` tool.

## Running the Application

Refer to the "Running the project" section for instructions on how to run the project using Docker Compose.

```bash
docker compose up
```

## Request examples

Request to create a new inventory item:

```bash
$ buf curl \
  --schema proto \
  --data '{"inventory": {"product_id": "123", "quantity": 225}}' \
  http://localhost:8080/inventory.v1.InventoryService/CreateInventory
```
