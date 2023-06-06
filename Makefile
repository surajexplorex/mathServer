generate_grpc_code:
	protoc protos/mathServer.proto --go-grpc_out=. --go_out=.

