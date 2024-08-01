package domain

// Cliente representa um cliente do restaurante
type Cliente struct {
	ID        string     `json:"id"`
	Nome      string     `json:"nome"`
	Telefone  string     `json:"telefone"`
	Email     string     `json:"email"`
	Feedbacks []Feedback `json:"feedbacks"` // Relacionamento: um cliente pode deixar muitos feedbacks
	Orders    []Order    `json"orders"`     // Relacionamento: um cliente pode fazer muitos pedidos
}

type Cliente struct {
	ID        int64  `json:"id"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	Telefone  string `json:"telefone"`
	Feedbacks string `json:"feedbacks"` // Campo feedbacks do tipo string
	Orders    string `json:"orders"`
}
