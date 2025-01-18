package main

import (
	"log"
	"ottodigital-be/prisma/db"
	"ottodigital-be/routes"

	"github.com/gin-gonic/gin"
)

// Initialize the Prisma client
var client = db.NewClient()

func main() {
	// Connect to the database
	if err := client.Connect(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer client.Disconnect()

	// Initialize the Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterBrandRoutes(r, client)
	routes.RegisterVoucherRoutes(r, client)
	routes.RegisterCustomerRoutes(r, client)
	routes.RegisterTransactionRoutes(r, client)

	// Start the server
	r.Run(":8080")
}
