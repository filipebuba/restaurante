package domain

// Funcionario representa um funcionário do restaurante
type Funcionario struct {
	ID      int
	Nome    string
	Cargo   string
	Salario float64
	Orders  []Order // Relacionamento: um funcionário pode atender muitos pedidos
}
