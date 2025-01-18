package routes

import (
	"ottodigital-be/controllers"
	"ottodigital-be/prisma/db"

	"github.com/gin-gonic/gin"
)

func RegisterVoucherRoutes(r *gin.Engine, client *db.PrismaClient) {
	controller := controllers.NewVoucherController(client)
	r.POST("/vouchers", controller.CreateVoucher)
	r.GET("/vouchers", controller.GetVouchers)
	r.GET("/vouchers/:id", controller.GetVoucherById)
	r.GET("/vouchers/brand/:id", controller.GetVouchersByBrandId)
}
