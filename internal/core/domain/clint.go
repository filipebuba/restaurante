package domain

// Cliente representa um cliente do restaurante
type Cliente struct {
	ID        string `json:"id,omitempty"`
	Nome      string `json:"nome"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
	Feedbacks string `json:"feedbacks"` // Relacionamento: um cliente pode deixar muitos feedbacks
	Orders    string `json:"orders"`    // Relacionamento: um cliente pode fazer muitos pedidos
}
