package main

import (
	"context"
	pb "github.com/pratyush934/mastering-gobackend/common/api"
	"google.golang.org/grpc"
	"log"
)

type GrpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGrPCHandler(grpcServer *grpc.Server) {
	handler := &GrpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *GrpcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New Order received...")
	p := &pb.Order{
		ID: "42",
	}
	return p, nil
}
