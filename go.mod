module github.com/ryan95z/grpc-to-rest-envoy-example/books-service

go 1.15

require (
	google.golang.org/genproto v0.0.0-20210603172842-58e84a565dcf // indirect
	google.golang.org/grpc v1.38.0
	v1 v1.0.0
)

replace ( 
	v1 v1.0.0 => ./books/v1
)