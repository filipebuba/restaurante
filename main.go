package main

import (
	"github.com/filipebuba/restaurante/internal/core/service"
	handlers "github.com/filipebuba/restaurante/internal/handler"
	mysqlrepo "github.com/filipebuba/restaurante/internal/repositories/mysql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	connect, err := mysqlrepo.GetConnectionDB()
	if err != nil {
		return
	}

	repo := mysqlrepo.NewMySQLRepository(connect)
	service := service.NewService(repo)

	h := handlers.NewHandler(service)

	// Endpoints para Cliente
	r.GET("/clientes", h.GetClientes)
	r.GET("/clientes/{id}", h.GetClienteByID)
	r.POST("/clientes", h.CreateCliente)
	r.PUT("/clientes/{id}", h.UpdateCliente)
	r.DELETE("/clientes/{id}", h.DeleteCliente)

	// Endpoints para Funcionário
	//r.GET("/funcionarios", getFuncionarios)
	//r.POST("/funcionarios", createFuncionario)

	// Endpoints para Menu
	//r.GET("/menus", getMenus)
	//r.POST("/menus", createMenu)

	// Endpoints para Food Item
	//r.GET("/food-items", getFoodItems)
	//r.POST("/food-items", createFoodItem)

	// Endpoints para Order
	//r.GET("/orders", getOrders)
	//r.POST("/orders", createOrder)

	// Endpoints para Invoice
	//r.GET("/invoices", getInvoices)
	//r.POST("/invoices", createInvoice)

	// Endpoints para Table
	//r.GET("/tables", getTables)
	//r.POST("/tables", createTable)

	// Endpoints para Promoção
	//r.GET("/promocoes", getPromocoes)
	//r.POST("/promocoes", createPromocao)

	// Endpoints para Feedback
	//r.GET("/feedbacks", getFeedbacks)
	//r.POST("/feedbacks", createFeedback)

	r.Run(":8080") // inicia o servidor na porta 8080
}
