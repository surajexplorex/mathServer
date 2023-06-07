generate_grpc_code:
	protoc proto/mathServer.proto --go-grpc_out=. --go_out=. --grpc-gateway_out=.

