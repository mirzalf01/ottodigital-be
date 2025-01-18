package controllers

import (
	"context"
	"net/http"
	"ottodigital-be/prisma/db"
	"ottodigital-be/responsehelper"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	client *db.PrismaClient
}

func NewTransactionController(client *db.PrismaClient) *TransactionController {
	return &TransactionController{client: client}
}

func (bc *TransactionController) CreateTransaction(c *gin.Context) {
	var transaction struct {
		Total      int      `json:"total"`
		CustomerId string   `json:"customerId"`
		VoucherIds []string `json:"voucherIds"`
	}

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, responsehelper.ResponseError("Something went wrong", err.Error()))
		return
	}

	_, err := bc.client.Customers.FindUnique(db.Customers.ID.Equals(transaction.CustomerId)).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, responsehelper.ResponseError("Customer not found", err.Error()))
		return
	}

	if len(transaction.VoucherIds) > 0 {
		// Find vouchers and check their validity
		vouchers, err := bc.client.Vouchers.FindMany(
			db.Vouchers.ID.In(transaction.VoucherIds),
			db.Vouchers.IsValid.Equals(true),
		).Exec(context.Background())

		if err != nil {
			c.JSON(http.StatusInternalServerError, responsehelper.ResponseError("Database query failed", err.Error()))
			return
		}

		existingVoucherIds := make(map[string]bool)
		for _, voucher := range vouchers {
			existingVoucherIds[voucher.ID] = true
		}

		// Find invalid voucher IDs (either not found or invalid)
		var invalidVoucherIds []string
		for _, voucherId := range transaction.VoucherIds {
			if _, exists := existingVoucherIds[voucherId]; !exists {
				invalidVoucherIds = append(invalidVoucherIds, voucherId)
			}
		}

		// If there are invalid voucher IDs, return them
		if len(invalidVoucherIds) > 0 {
			c.JSON(http.StatusNotFound, responsehelper.ResponseError("Invalid voucher IDs", invalidVoucherIds))
			return
		}
	}

	// Proceed to create the transaction
	var transactionData *db.TransactionsModel

	if len(transaction.VoucherIds) > 0 {
		transactionData, err = bc.client.Transactions.CreateOne(
			db.Transactions.Customer.Link(db.Customers.ID.Equals(transaction.CustomerId)),
			db.Transactions.Total.Set(transaction.Total),
		).Exec(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, responsehelper.ResponseError("Something went wrong", err.Error()))
		}

		// Invalidate the vouchers once the transaction is successful
		_, err = bc.client.Vouchers.FindMany(db.Vouchers.ID.In(transaction.VoucherIds)).
			Update(db.Vouchers.IsValid.Set(false), db.Vouchers.TransactionID.Set(transactionData.ID)).
			Exec(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, responsehelper.ResponseError("Something went wrong", err.Error()))
			return
		}
	} else {
		transactionData, err = bc.client.Transactions.CreateOne(
			db.Transactions.Customer.Link(db.Customers.ID.Equals(transaction.CustomerId)),
			db.Transactions.Total.Set(transaction.Total),
		).Exec(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, responsehelper.ResponseError("Something went wrong", err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, responsehelper.ResponseSuccess("Success", transactionData))
}

func (bc *TransactionController) GetTransactions(c *gin.Context) {
	transactions, err := bc.client.Transactions.FindMany().With(
		db.Transactions.Customer.Fetch(),
		db.Transactions.Vouchers.Fetch().With(db.Vouchers.Brand.Fetch()),
	).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responsehelper.ResponseError("Something went wrong", err.Error()))
		return
	}

	c.JSON(http.StatusOK, responsehelper.ResponseSuccess("Success", transactions))
}

func (bc *TransactionController) GetTransactionById(c *gin.Context) {
	id := c.Param("id")

	transaction, err := bc.client.Transactions.FindUnique(db.Transactions.ID.Equals(id)).With(db.Transactions.Vouchers.Fetch().With(db.Vouchers.Brand.Fetch())).Exec(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, responsehelper.ResponseError("Transaction not found", err.Error()))
		return
	}

	c.JSON(http.StatusOK, responsehelper.ResponseSuccess("Success", transaction))
}
