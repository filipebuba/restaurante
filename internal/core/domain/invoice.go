package domain

// Invoice representa uma fatura de um pedido
type Invoice struct {
	ID              int
	MétodoPagamento string
	StatusPagamento string
	//Order           Order // Relacionamento: uma fatura é gerada para um pedido
}
