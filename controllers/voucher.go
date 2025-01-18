package controllers

import (
	"context"
	"net/http"
	"ottodigital-be/prisma/db"
	"ottodigital-be/responsehelper"

	"github.com/gin-gonic/gin"
)

type VoucherController struct {
	client *db.PrismaClient
}

func NewVoucherController(client *db.PrismaClient) *VoucherController {
	return &VoucherController{client: client}
}

// Implement methods like CreateVoucher, GetVouchers, etc.
func (bc *VoucherController) CreateVoucher(c *gin.Context) {
	var voucher struct {
		Name    string `json:"name" binding:"required"`
		Point   int    `json:"point" binding:"required,min=1"`
		BrandId string `json:"brandId" binding:"required"`
	}

	// Validate JSON input
	if err := c.ShouldBindJSON(&voucher); err != nil {
		c.JSON(http.StatusBadRequest, responsehelper.ResponseError("Invalid input", err.Error()))
		return
	}

	// Check if brand exists
	brand, err := bc.client.Brands.FindUnique(db.Brands.ID.Equals(voucher.BrandId)).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, responsehelper.ResponseError("Brand not found", err.Error()))
		return
	}

	newVoucher, err := bc.client.Vouchers.CreateOne(
		db.Vouchers.Point.Set(int(voucher.Point)),
		db.Vouchers.Brand.Link(db.Brands.ID.Equals(brand.ID)),
		db.Vouchers.Name.Set(voucher.Name),
	).Exec(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, responsehelper.ResponseError("Failed to create voucher", err.Error()))
		return
	}

	c.JSON(http.StatusOK, responsehelper.ResponseSuccess("Voucher created successfully", newVoucher))
}

func (bc *VoucherController) GetVouchers(c *gin.Context) {
	vouchers, err := bc.client.Vouchers.FindMany().With(db.Vouchers.Brand.Fetch()).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responsehelper.ResponseError("Something went wrong", err.Error()))
		return
	}

	c.JSON(http.StatusOK, responsehelper.ResponseSuccess("Success", vouchers))
}

func (bc *VoucherController) GetVoucherById(c *gin.Context) {
	id := c.Param("id")

	voucher, err := bc.client.Vouchers.FindUnique(db.Vouchers.ID.Equals(id)).With(db.Vouchers.Brand.Fetch()).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, responsehelper.ResponseError("Voucher not found", err.Error()))
		return
	}

	c.JSON(http.StatusOK, responsehelper.ResponseSuccess("Success", voucher))
}

func (bc *VoucherController) GetVouchersByBrandId(c *gin.Context) {
	id := c.Param("id")

	brand, err := bc.client.Brands.FindUnique(db.Brands.ID.Equals(id)).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, responsehelper.ResponseError("Brand not found", err.Error()))
		return
	}

	vouchers, err := bc.client.Vouchers.FindMany(
		db.Vouchers.Brand.Where(
			db.Brands.ID.Equals(id),
		),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responsehelper.ResponseError("Something went wrong", err.Error()))
		return
	}

	data := gin.H{
		"brand":    brand,
		"vouchers": vouchers,
	}

	c.JSON(http.StatusOK, responsehelper.ResponseSuccess("Success", data))
}
