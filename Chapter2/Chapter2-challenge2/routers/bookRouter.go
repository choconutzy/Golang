package routers

import (
	"Chapter2-challenge2/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome",
		})
	})
	router.POST("/book", controllers.CreateBook)
	router.GET("/book/:id", controllers.GetBookById)
	router.GET("/books", controllers.GetBooks)
	router.PATCH("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)

	return router
}
