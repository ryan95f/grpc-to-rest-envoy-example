# Golang gRPC to REST Envoy example

## Requirements

The following are required:

* Docker
* Golang
* protoc (Protocal Buffer compiler)

## Compiling the protocal buffers

This will compile the `books.proto` protocol buffer and create the descriptor binary required for the envoy proxy.

```bash
make protoc
```

## Creating the Docker containers

```bash
docker-compose build
```

## Running the Docker containers

```bash
docker-compose up -d
```

## Validating the envoy config

```bash
docker run --rm \
       -v $(pwd)/envoy/envoy.yaml:/envoy.yaml envoyproxy/envoy:v1.17-latest \
       --mode validate -c envoy.yaml
```
