package routes

import (
	"ottodigital-be/controllers"
	"ottodigital-be/prisma/db"

	"github.com/gin-gonic/gin"
)

func RegisterBrandRoutes(r *gin.Engine, client *db.PrismaClient) {
	controller := controllers.NewBrandController(client)
	r.POST("/brands", controller.CreateBrand)
	r.GET("/brands", controller.GetBrands)
}
