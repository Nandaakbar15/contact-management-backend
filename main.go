package main

import (
	"contact-management-backend/models"
	"contact-management-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()
	r.Use(cors.Default())
	routes.Route(r)

	r.Run() // Default port is :8080
}
