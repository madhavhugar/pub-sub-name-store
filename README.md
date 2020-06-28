# Pub-Sub Pattern

Implementation of a simple topic based pub-sub arhitecture with a producer(in Go) & a consumer(in Python). The communication between the producer and consumer is handled by a message broker(rabbitMQ) using AMQP protocol.

### Architecture in a single line:

csv-data => producer => rabbitMQ => consumer => postgres

### TODO

- [] Search all TODO statements
- [] Proper error handling