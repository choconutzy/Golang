package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var Books = []Book{
	{
		BookID: 1,
		Title:  "Haruna kagawa",
		Author: "Nakamura",
		Desc:   "evudb8ureuferu",
	},
	{
		BookID: 2,
		Title:  "Haruna kagawa",
		Author: "Nakamura",
		Desc:   "evudb8ureuferu",
	},
	{
		BookID: 3,
		Title:  "Haruna kagawa",
		Author: "Nakamura",
		Desc:   "evudb8ureuferu",
	},
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = len(Books) + 1

	Books = append(Books, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "book has been successfully created",
	})
}

func GetAllBook(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"book": Books,
	})
}

func GetBookById(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false
	var bookData Book

	for i, book := range Books {
		id := fmt.Sprintf("%v", book.BookID)
		fmt.Print(id)
		if bookID == (id) {
			condition = true
			bookData = Books[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range Books {
		id := fmt.Sprintf("%v", book.BookID)
		fmt.Print(id)
		if bookID == (id) {
			condition = true

			Books[i] = updatedBook
			Books[i].BookID = book.BookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully updated", bookID),
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("id")
	condition := false
	var bookIndex int

	for i, book := range Books {
		id := fmt.Sprintf("%v", book.BookID)
		fmt.Print(id)
		if bookID == (id) {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	copy(Books[bookIndex:], Books[bookIndex-1:])
	Books = Books[:len(Books)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully deleted", bookID),
	})
}
