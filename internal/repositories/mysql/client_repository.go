package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/filipebuba/restaurante/internal/core/domain"
	"github.com/filipebuba/restaurante/internal/core/ports"
	"github.com/jmoiron/sqlx"
)

type clientRepositoryImpl struct {
	connect *sqlx.DB
}

func NewMySQLRepository(connect *sqlx.DB) ports.ClienteRepository {
	return &clientRepositoryImpl{
		connect: connect,
	}
}

func (r *clientRepositoryImpl) GetAllClientes(ctx context.Context) ([]domain.Cliente, error) {
	var client []domain.Cliente

	err := r.connect.Select(&client, "SELECT * FROM clients LIMIT 20")
	if err != nil {
		return client, fmt.Errorf("error getting all client: %w", err)
	}

	return client, nil
}

func (r *clientRepositoryImpl) CreateCliente(ctx context.Context, client domain.Cliente) (*domain.Cliente, error) {
	result, err := r.connect.Exec("INSERT INTO clients (nome, email, telefone, feedbacks, orders) VALUES (?,?,?,?,?)", client.Nome, client.Email, client.Telefone, client.Feedbacks, client.Orders)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("error creating client: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error saving client: %w", err)
	}

	client.ID = strconv.FormatInt(id, 20)

	return &client, nil
}

func (r *clientRepositoryImpl) UpdateCliente(ctx context.Context, editCliente domain.Cliente) (*domain.Cliente, error) {
    query := `
        UPDATE clients 
        SET nome = ?, email = ?, telefone = ?, feedbacks = ?, orders = ? 
        WHERE id = ?
    `
    _, err := r.connect.ExecContext(ctx, query, 
        editCliente.Nome, 
        editCliente.Email, 
        editCliente.Telefone, 
        editCliente.Feedbacks, 
        editCliente.Orders, 
        editCliente.ID,
    )
    if err != nil {
        return nil, fmt.Errorf("error updating client: %w", err)
    }

    return &editCliente, nil
}

func (r *clientRepositoryImpl) DeleteCliente(ctx context.Context, id string) error {
	_, err := r.connect.ExecContext(ctx, "DELETE FROM clients WHERE id=?", id)
	if err != nil {
		return fmt.Errorf("error deleting client: %w", err)
	}

	return nil
}

func (r *clientRepositoryImpl) GetClienteByID(ctx context.Context, id string) (*domain.Cliente, error) {
	var client domain.Cliente

	err := r.connect.Get(&client, "SELECT * FROM clients WHERE id=?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("client not found")
		}
		return nil, fmt.Errorf("error getting client by ID: %w", err)
	}

	return &client, nil
}