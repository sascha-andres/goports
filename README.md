# goports

... is a library implementing ports for communication with the outside for microservices.

The basic idea is, that you get data in through an incoming port and let data out throught an outgoing port. *Only* through ports. Only possible exceptions I can think of are stuff like prometheus endpoints, etc

## Why

Because this abstracts away the need for some low level work and forces you into a cleaner architecture.

## Architecture

It is based on two interfaces (IncomingPort anbd OutgoingPort (sursprise, surprise)) with implementations for different transportation layers (like RabbitMQ or provding a HTTP endpoint, etc)

All configuration is done using environment variables.