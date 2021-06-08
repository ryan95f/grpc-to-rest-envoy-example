# gRPC to REST Envoy example

An example of how to use the envoy proxy to transcode REST calls to gRPC calls and expose the gRPC service on a common port

## Requirements

The following are required:

* Docker
* Golang 15 or greater
* protoc (Protocal Buffer compiler)

## Compiling the protocal buffers

This will compile the `books.proto` protocol buffer and create the descriptor binary required for the envoy proxy.

```bash
make protoc
```

## Creating the Docker images

To build the Docker containers, run:

```bash
docker-compose build
```

This will create two images:

* `books-envoy` which is the envoy proxy
* `books` which is the book service

To build a specific image, run:

```bash
docker-compose build <service-name>
```

Where `<service-name>` can be replaced with `envoy` for the envoy image or `books` for the book service.

## Running the Docker containers

To create containers for both `books` and `envoy`, run:

```bash
docker-compose up -d
```

## Validating the Envoy config

To validate that your Envoy config is valid, run:

```bash
docker run --rm \
       -v $(pwd)/envoy/envoy.yaml:/envoy.yaml envoyproxy/envoy:v1.17-latest \
       --mode validate -c envoy.yaml
```
