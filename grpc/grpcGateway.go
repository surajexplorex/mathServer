package grpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	gw "mathOperation/proto/mathServer"
	"net/http"
)

func InitGRPCGateway() {

	mux := runtime.NewServeMux()
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	//defer cancel()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterMathOperationsHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC gateway service endpoint: %v", err)
	}

	go func() {
		if err := http.ListenAndServe(":8081", mux); err != nil {
			log.Fatalf("Could not setup HTTP endpoint: %v", err)
		}
	}()
	log.Printf("============ GRPC Gateway started.============ ")
}
