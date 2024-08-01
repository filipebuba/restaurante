package domain

// Feedback representa o feedback de um cliente sobre um pedido
type Feedback struct {
	ID         int
	Comentario string
	Avaliacao  int
	Cliente    Cliente // Relacionamento: um feedback é deixado por um cliente
	//Order      Order   // Relacionamento: um feedback refere-se a um pedido
}
