package handlers

// Cliente representa um cliente do restaurante
type cliente struct {
	ID        string `json:"id"`
	Nome      string `json:"nome" validate:"required"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email" validate:"required"`
	Feedbacks string `json:"feedbacks"` // Relacionamento: um cliente pode deixar muitos feedbacks
	Orders    string `json"orders"`     // Relacionamento: um cliente pode fazer muitos pedidos
}
