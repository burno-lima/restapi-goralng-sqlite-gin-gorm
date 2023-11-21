package main

import (
	"github.com/gin-gonic/gin"

	"example.com/RestAPIgo/connector"
	"example.com/RestAPIgo/controllers"
)

func main() {

	r := gin.Default()

	connector.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	err := r.Run()
	if err != nil {
		return
	}
}
