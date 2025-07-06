package main

import (
	"github.com/pratyush934/mastering-gobackend/common"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {

	server := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)

	if err != nil {
		log.Fatal(err.Error())
	}

	//store := NewStore()
	//service := NewService(store)
	NewGrPCHandler(server)

	//_ = service.CreateOrder(context.Background())

	log.Println("GRPC Server started ", grpcAddr)

	if err := server.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
