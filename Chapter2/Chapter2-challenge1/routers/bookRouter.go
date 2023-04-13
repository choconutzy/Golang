package routers

import (
	"Chapter2-challenge1/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/book", controllers.CreateBook)
	router.GET("/book/:id", controllers.GetBookById)
	router.GET("/books", controllers.GetAllBook)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)

	return router
}
