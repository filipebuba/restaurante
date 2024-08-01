package domain

// Menu representa o menu do restaurante
type Menu struct {
	ID        int
	Nome      string
	Categoria string
	FoodItems []FoodItem // Relacionamento: um menu pode conter muitos itens de comida
}
