package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	protos "mathOperation/proto/mathServer"
	"mathOperation/service"
	"net"
)

func StartGRPCServer() {

	// create new gRPC server
	s := grpc.NewServer()

	// create new instance of MathOperation grpc
	trans := service.NewMathOperationServer()

	// register reflection API
	reflection.Register(s)

	// register it to the grpc
	protos.RegisterMathOperationsServer(s, trans)

	// create socket to listen to requests
	tl, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(fmt.Println("Error starting tcp listener on port 8080", err))
	}
	fmt.Println("starting tcp listener on port 8080")
	log.Printf("============ GRPC server started.============ ")

	// start listening
	s.Serve(tl)

}
