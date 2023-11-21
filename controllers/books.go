package controllers

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
	"example.com/RestAPIgo/models"
	"example.com/RestAPIgo/connector"
)

func FindBooks(c *gin.Context){
	var books []models.Book
	connector.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	book := models.Book{Title: input.Title, Author: input.Author}
	connector.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book models.Book

	if err := connector.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
  // Get model if exist
  var book models.Book

	if err := connector.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  // Validate input
  var input models.UpdateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  connector.DB.Model(book).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := connector.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	connector.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
