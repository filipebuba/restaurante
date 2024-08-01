package ports

import (
	"context"

	"github.com/filipebuba/restaurante/internal/core/domain"
)

type ClienteRepository interface {
	GetAllClientes(ctx context.Context, limit int, searchAfter []interface{}) ([]domain.Cliente, interface{}, error)
	CreateCliente(ctx context.Context, client domain.Cliente) (*domain.Cliente, error)
	UpdateCliente(ctx context.Context, editCliente domain.Cliente) (*domain.Cliente, error)
	DeleteCliente(ctx context.Context, id string) error
}
