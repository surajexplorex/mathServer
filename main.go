package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"mathOperation/db"
	protos "mathOperation/protos/mathServer"
	"mathOperation/server"
	"net"
)

func main() {

	// initializing the database
	err := db.InitTestDB()
	if err != nil {
		log.Fatal(fmt.Println("Error starting the DB", err))
	}

	// starting the GRPC server
	server := CreateGRPCSever()

	// create socket to listen to requests
	tl, err := net.Listen("tcp", "localhost:8765")
	if err != nil {
		log.Fatal(fmt.Println("Error starting tcp listener on port 8765", err))
	}
	fmt.Println("starting tcp listener on port 8765")

	// start listening
	server.Serve(tl)

}

func CreateGRPCSever() *grpc.Server {
	// create new gRPC server
	s := grpc.NewServer()

	// create new instance of MathOperation server
	trans := server.NewMathOperationServer()

	// register reflection API https://github.com/grpc/grpc/blob/master/doc/server-reflection.md
	reflection.Register(s)

	// register it to the grpc server
	protos.RegisterMathOperationsServer(s, trans)

	return s

}
