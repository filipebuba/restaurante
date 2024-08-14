package handlers

import "github.com/filipebuba/restaurante/internal/core/ports"

type handler struct {
	clientService ports.ClienteService
}

func NewHandler(clientService ports.ClienteService) *handler {
	return &handler{
		clientService: clientService,
	}
}