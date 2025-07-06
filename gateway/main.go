package main

import (
	pb "github.com/pratyush934/mastering-gobackend/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	// common "github.com/pratyush934/mastering-gobackend/common"
)

var ordersServiceAdd = "localhost:3000"

func main() {

	conn, err2 := grpc.NewClient(ordersServiceAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err2 != nil {
		log.Fatalln("In the main.go/gateway service ", err2)
	}

	client := pb.NewOrderServiceClient(conn)

	err := godotenv.Load("gateway/.env")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			log.Println("Warning: Could not load .env file:", err)
		}
	}
	httpAddr := os.Getenv("HTTP_ADDR")

	if httpAddr == "" {
		httpAddr = ":5050"
	}

	log.Printf("HTTP_ADDR environment variable: %s", httpAddr)

	mux := http.NewServeMux()
	handler := NewHandler(client)
	handler.RegisterRoutes(mux)

	log.Println("It is just working and you should be proud of that", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start", err)
	}
}
