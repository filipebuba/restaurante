package domain

// Cliente representa um cliente do restaurante
type Cliente struct {
	ID        string 
	Nome      string 
	Telefone  string 
	Email     string 
	Feedbacks string  // Relacionamento: um cliente pode deixar muitos feedbacks
	Orders    string    // Relacionamento: um cliente pode fazer muitos pedidos
}
