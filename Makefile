GO := go
APP_NAME := books-service
PROTOC_DESCRIPTOR := protos/books.pb
GOOGLE_APIS_DIR := ../googleapis

edit: build run

build:
	@${GO} build -o ${APP_NAME}

run:
	@${APP_NAME}

protoc:
	@protoc -I./${GOOGLE_APIS_DIR} -I. \
		--include_imports \
		--include_source_info \
    	--descriptor_set_out=${PROTOC_DESCRIPTOR} \
		--go-grpc_out=. \
		--go_out=. protos/books.proto

clean:
	@rm ${APP_NAME}
	@rm ${PROTOC_DESCRIPTOR}