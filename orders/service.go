package main

import (
	"context"
	"log"
)

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service {
	return &service{store: store}
}

func (s *service) CreateOrder(context.Context) error {
	log.Println("Hello I am I am")
	return nil
}
