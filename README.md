# gRPC to REST Envoy example

An example of how to use the envoy proxy to transcode REST requests to gRPC calls and expose the gRPC service on a single port.

## Requirements

The following are required:

* Docker
* Golang 15 or greater
* protoc (Protocal Buffer compiler)

## Set Up

In order to compile the `books.proto` file, you first need to clone the [Google API](https://github.com/googleapis/googleapis) repository from Github. This will include the `annotations.proto` file, which is required to provide the capabilities to transcode between REST and gRPC. Then you need to add the Google repository to your local path:

```bash
git clone https://github.com/googleapis/googleapis

GOOGLE_APIS_DIR = <local-google-api-location>
```

## Compiling the protocal buffers

In order to compile the `books.proto` file, you need to run the `protoc` compiler. This can be done with the included `Makefile`:

```bash
make protoc
```

Alternatively, you can directly run the `protoc` compiler:

```bash
protoc -I./${GOOGLE_APIS_DIR} -I. \
       --include_imports \
       --include_source_info \
       --descriptor_set_out=protos/books.pb \
       --go-grpc_out=. \
       --go_out=. protos/books.proto
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

This will run the API on port `8080` and the envoy admin view on port `9901`.

## Running the Books service locally

In order to run the books service locally, first build the application:

```bash
make build
```

Then run the application. By default it will be on port `9000`:

```bash
make run
```

## Deploying on Kubernetes

The books service with envoy can be deployed with Kubernetes. The following files are provided:

* `books-configmap.yaml` - Defines a config map that contains the envoy config for running the service on Kubernetes.
* `books.yaml` - Defines the deployment and service for running the application. The service uses a `NodePort` and will expose the deployment on port `30001`, which can accept REST and gRPC calls.

To deploy the config map:

```bash
kubectl apply -f books-configmap.yaml
```

Then to deploy the service and deployment:

```bash
kubectl apply -f books.yaml
```

**Note**: By default the number of replicas in the deployment is set to 1. Since the data is stored within memory on the container and not in a database, scaling the deployment above 1 may result in differences in the responses with each request.

## Validating the Envoy config

To validate that your Envoy config is valid, run:

```bash
docker run --rm \
       -v $(pwd)/envoy.yaml:/envoy.yaml envoyproxy/envoy:v1.17-latest \
       --mode validate -c envoy.yaml
```
