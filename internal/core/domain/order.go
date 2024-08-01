package domain

import "time"

// Order representa um pedido feito por um cliente
type Order struct {
	ID           int
	DataPedido   time.Time
	StatusPedido string
	Cliente      Cliente     // Relacionamento: um pedido é feito por um cliente
	Funcionario  Funcionario // Relacionamento: um pedido é atendido por um funcionário
	Invoice      Invoice     // Relacionamento: um pedido gera uma fatura
	OrderItems   []OrderItem // Relacionamento: um pedido contém muitos itens de pedido
	Table        Table       // Relacionamento: um pedido está associado a uma mesa
	Feedback     Feedback    // Relacionamento: um pedido pode ter um feedback
}
