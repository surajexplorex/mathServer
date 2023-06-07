package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	protos "mathOperation/proto/mathServer"
	"net"
)

func StartGRPCServer() {

	// create new gRPC server
	s := grpc.NewServer()

	// create new instance of MathOperation grpc
	trans := NewMathOperationServer()

	// register reflection API https://github.com/grpc/grpc/blob/master/doc/server-reflection.md
	reflection.Register(s)

	// register it to the grpc grpc
	protos.RegisterMathOperationsServer(s, trans)

	// create socket to listen to requests
	tl, err := net.Listen("tcp", "localhost:8765")
	if err != nil {
		log.Fatal(fmt.Println("Error starting tcp listener on port 8765", err))
	}
	fmt.Println("starting tcp listener on port 8765")

	// start listening
	s.Serve(tl)

}
