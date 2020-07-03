# Pub-Sub Pattern

Implementation of a simple topic based pub-sub arhitecture with a producer(in Go) & a consumer(in Python). The communication between the producer and consumer is handled by a message broker(rabbitMQ) using AMQP protocol.

### Architecture in a single line:

```
csv-data => producer(golang) => rabbitMQ => consumer(python) => postgres
```

### Setup

```
docker-compose up
```

### Execute tests

```
cd producer
go test ./...
```

```
cd consumer
python -m pytest tests/
```

### TODO

- Test functionality that depends on rabbitMQ - still quite new to golang testing ecosystem
- Ensure services are started within docker-compose, rather then sleeping within application
- Add logging