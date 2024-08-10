package main

import (
	"ukashanoor/event-booking/db"
	"ukashanoor/event-booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.InitDB()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
