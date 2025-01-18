package controllers

import (
	"context"
	"net/http"
	"ottodigital-be/prisma/db"
	"ottodigital-be/responsehelper"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	client *db.PrismaClient
}

func NewCustomerController(client *db.PrismaClient) *CustomerController {
	return &CustomerController{client: client}
}

func (bc *CustomerController) CreateCustomer(c *gin.Context) {
	var customer struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, responsehelper.ResponseError("Something went wrong", err.Error()))
		return
	}

	newCustomer, err := bc.client.Customers.CreateOne(
		db.Customers.Name.Set(customer.Name),
		db.Customers.Address.Set(customer.Address),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, responsehelper.ResponseError("Something went wrong", err.Error()))
		return
	}
	c.JSON(http.StatusOK, responsehelper.ResponseSuccess("Success", newCustomer))
}

func (bc *CustomerController) GetCustomers(c *gin.Context) {
	customers, err := bc.client.Customers.FindMany().Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responsehelper.ResponseError("Something went wrong", err.Error()))
		return
	}

	c.JSON(http.StatusOK, responsehelper.ResponseSuccess("Success", customers))
}
