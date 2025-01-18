package routes

import (
	"ottodigital-be/controllers"
	"ottodigital-be/prisma/db"

	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoutes(r *gin.Engine, client *db.PrismaClient) {
	controller := controllers.NewTransactionController(client)
	r.POST("/transaction/redemption", controller.CreateTransaction)
	r.GET("/transaction/redemption", controller.GetTransactions)
	r.GET("/transaction/redemption/:id", controller.GetTransactionById)
}
