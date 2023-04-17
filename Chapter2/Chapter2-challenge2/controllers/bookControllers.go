package controllers

import (
	"Chapter2-challenge2/config"
	"Chapter2-challenge2/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetBooks(ctx *gin.Context) {
	var books []models.Book
	config.DB.Find(&books)
	message := "Successfully get all books"
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    books,
	})
}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("id")
	var bookData models.Book
	find := config.DB.First(&bookData, "id = ?", bookID)
	if find.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": fmt.Sprint("Data not found for book with id ", bookID),
			"error":   find.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func CreateBook(ctx *gin.Context) {
	var newBook models.Book
	newBook.ID = uuid.New()
	json.Marshal(newBook)
	ctx.ShouldBindJSON(&newBook)
	data := config.DB.Create(&newBook)

	if data.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add book",
			"error":   data.Error.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "book has been successfully created",
		"data":    newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	var bookData models.Book
	find := config.DB.Where("id = ?", bookID).First(&bookData)
	if find.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": find.Error.Error()})
		return
	}

	ctx.BindJSON(&bookData)
	config.DB.Model(&bookData).Updates(bookData)
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %s successfully updated", bookID),
		"data":    bookData,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	var bookData models.Book
	find := config.DB.Where("id = ?", bookID).First(&bookData)

	if find.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": find.Error.Error()})
		return
	}

	config.DB.Delete(bookData)

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %s successfully deleted", bookID),
	})
}
