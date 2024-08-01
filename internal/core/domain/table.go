package domain

// Table representa uma mesa no restaurante
type Table struct {
	ID         int
	NumeroMesa int
	Orders     []Order // Relacionamento: uma mesa pode estar associada a muitos pedidos
}
