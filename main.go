package main

import (
	"github.com/ChubbyJoe/bloggr/models"
	"github.com/ChubbyJoe/bloggr/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// This function connects to the database when the app starts
	models.ConnectDB()

	r := gin.Default()

	// This function registers routes
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
