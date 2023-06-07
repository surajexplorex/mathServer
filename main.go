package main

import (
	"fmt"
	"log"
	"mathOperation/db"
	"mathOperation/grpc"
)

func main() {

	// initializing the database
	err := db.InitTestDB()
	if err != nil {
		log.Fatal(fmt.Println("Error starting the DB", err))
	}

	// startGRPC Gateway
	grpc.InitGRPCGateway()

	// startGRPC Server
	grpc.StartGRPCServer()

}
