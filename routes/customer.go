package routes

import (
	"ottodigital-be/controllers"
	"ottodigital-be/prisma/db"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(r *gin.Engine, client *db.PrismaClient) {
	controller := controllers.NewCustomerController(client)
	r.POST("/customers", controller.CreateCustomer)
	r.GET("/customers", controller.GetCustomers)
}
