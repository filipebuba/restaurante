package domain

// FoodItem representa um item de comida no menu
type FoodItem struct {
	ID         int
	Nome       string
	Preco      float64
	Menus      []Menu      // Relacionamento: um item de comida pode estar em muitos menus
	OrderItems []OrderItem // Relacionamento: um item de comida pode aparecer em muitos pedidos
	Promocoes  []Promocao  // Relacionamento: um item de comida pode ter muitas promoções
}
