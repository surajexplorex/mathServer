package gateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	gw "mathOperation/proto/mathServer"
	"net/http"
)

const (
	grpcServerEndpoint = "localhost:%s"
)

func InitGRPCGateway() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterMathOperationsHandlerFromEndpoint(ctx, mux, fmt.Sprintf(grpcServerEndpoint, "5270"), opts)
	if err != nil {
		log.Fatalf("Failed to register gRPC gateway service endpoint: %v", err)
	}
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("Could not setup HTTP endpoint: %v", err)
	}
}
