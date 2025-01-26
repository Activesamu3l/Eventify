package main

import (
	"github.com/Activesamu3l/Eventify/db"
	"github.com/Activesamu3l/Eventify/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}


