package domain

// OrderItem representa um item de um pedido
type OrderItem struct {
	ID            int
	Qtd           int
	PrecoUnitario float64
	Order         Order    // Relacionamento: um item de pedido pertence a um pedido
	FoodItem      FoodItem // Relacionamento: um item de pedido refere-se a um item de comida
}
