package main

import (
	"log"

	"github.com/chonlawit-odds/task-api/internal/item"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database
	dbconn, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/task"))
	if err != nil {
		log.Fatal(err)
	}

	// Controller
	controller := item.NewController(dbconn)

	// Server
	r := gin.Default()

	// Router
	r.POST("/api/v1/tasks", controller.CreateItem)
	r.GET("/api/v1/tasks", controller.GetItems)
	r.PUT("/api/v1/tasks/:id", controller.ReplaceItem)
	r.PATCH("/api/v1/tasks/:id", controller.UpdateItemStatus)
	r.DELETE("/api/v1/tasks/:id", controller.DeleteItem)

	// Listen port 8080 by default.
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
