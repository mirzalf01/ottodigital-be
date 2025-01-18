package controllers

import (
	"context"
	"net/http"
	"ottodigital-be/prisma/db"
	"ottodigital-be/responsehelper"

	"github.com/gin-gonic/gin"
)

type BrandController struct {
	client *db.PrismaClient
}

func NewBrandController(client *db.PrismaClient) *BrandController {
	return &BrandController{client: client}
}

func (bc *BrandController) CreateBrand(c *gin.Context) {
	var brand struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, responsehelper.ResponseError("Invalid input", err.Error()))
		return
	}

	newBrand, err := bc.client.Brands.CreateOne(
		db.Brands.Name.Set(brand.Name),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responsehelper.ResponseError("Failed to create brand", err.Error()))
		return
	}

	c.JSON(http.StatusOK, responsehelper.ResponseSuccess("Brand created successfully", newBrand))
}

func (bc *BrandController) GetBrands(c *gin.Context) {
	brands, err := bc.client.Brands.FindMany().Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responsehelper.ResponseError("Failed to fetch brands", err.Error()))
		return
	}

	c.JSON(http.StatusOK, responsehelper.ResponseSuccess("Success", brands))
}
