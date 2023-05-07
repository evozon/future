package collector

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() (CollectorClient, *grpc.ClientConn) {
	grpcConn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to grpc server: %v\n", err)
	}

	return NewCollectorClient(grpcConn), grpcConn
}
