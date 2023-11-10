proto:
	protoc -I./protos   --go_out=./protos   --go-grpc_out=./protos \
	--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
 	./protos/rusprofile_lookup.proto

swagger:
	protoc -I./protos --openapiv2_out=./protos \
    --openapiv2_opt=logtostderr=true \
    ./protos/rusprofile_lookup.proto


docker:
	docker build -t inn-lookup-pro .
	docker build -t my-swagger-ui .
