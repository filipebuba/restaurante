package domain

import "time"

// Promocao representa uma promoção aplicada a itens do menu
type Promocao struct {
	ID                 int
	Descricao          string
	PercentualDesconto float64
	DataValidade       time.Time
	FoodItems          []FoodItem // Relacionamento: uma promoção pode se aplicar a muitos itens de comida
}
