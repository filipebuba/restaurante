package service

import (
	"context"
	"fmt"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/filipebuba/restaurante/internal/core/ports"
)

type clientServiceImpl struct {
	repo ports.ClienteRepository
}

func NewService(repo ports.ClienteRepository) ports.ClienteService {
	return clientServiceImpl{
		repo: repo,
	}
}

func (s clientServiceImpl) GetAllClientes(ctx context.Context, limit int, cursor []interface{}) ([]domain.Cliente, interface{}, error) {
	return s.repo.GetAllClientes(ctx, limit, cursor)
}

func (s clientServiceImpl) CreateCliente(ctx context.Context, client domain.Cliente) (*domain.Cliente, error) {
	if client.Nome == "" {
		return nil, fmt.Errorf("name cannot be nil")
	}

	if client.Email == "" {
		return nil, fmt.Errorf("email cannot be nil")
	}

	clientNew, err := s.repo.CreateCliente(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("error in repository: %w", err)
	}

	return clientNew, nil
}

func (s clientServiceImpl) UpdateCliente(ctx context.Context, editCliente domain.Cliente) (*domain.Cliente, error) {
	if editCliente.Nome == "" {
		return nil, fmt.Errorf("name cannot be nil")
	}

	if editCliente.Email == "" {
		return nil, fmt.Errorf("email cannot be nil")
	}

	clientUpdated, err := s.repo.UpdateCliente(ctx, editCliente)
	if err != nil {
		return nil, fmt.Errorf("error in repository: %w", err)
	}

	return clientUpdated, nil
}

func (s clientServiceImpl) DeleteCliente(ctx context.Context, id string) error {
	err := s.repo.DeleteCliente(ctx, id)
	if err != nil {
		return fmt.Errorf("error in repository: %w", err)
	}
	return nil
}
