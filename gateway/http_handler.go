package main

import (
	pb "github.com/pratyush934/mastering-gobackend/common/api"
	"net/http"
)

type handler struct {
	client pb.OrderServiceClient
	//gateway
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client: client}
}

func (h *handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.HadleCreateOrder)
}

func (h *handler) HadleCreateOrder(w http.ResponseWriter, r *http.Request) {

	customerId := r.PathValue("customerId")

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerId: customerId,
	})

}
