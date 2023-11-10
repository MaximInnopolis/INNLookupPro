proto:
	protoc -I./protos   --go_out=./protos   --go-grpc_out=./protos \
	--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
 	./protos/rusprofile_lookup.proto