package controller

import (
	"LibraryAPI/config"
	"LibraryAPI/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuthors(c *gin.Context) {
	var author []model.Author
	var response model.AuthorResponse
	id := c.Query("id")
	db := config.Connect()

	if id == "" {
		if err := db.Find(&author).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := db.First(&author, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
			return
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = author
	c.JSON(http.StatusOK, response)
}

func CreateAuthor(c *gin.Context) {
	var author model.Author
	var response model.Response

	if err := c.BindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.Connect()

	if err := db.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Status = 201
	response.Message = "Insert Author Success"
	c.JSON(http.StatusCreated, response)
}

func UpdateAuthor(c *gin.Context) {
	var author model.Author
	var response model.Response
	db := config.Connect()

	//จาก JSON -> author
	if err := c.BindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(author)
	//เช็คว่ามีตัวที่แก้ไขไหมจาก id
	if err := db.First(&model.Author{}, author.AUTHOR_ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	//บันทึกการแก้ไข
	fmt.Println(author)
	if err := db.Save(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Status = http.StatusOK
	response.Message = "Author updated successfully"

	c.JSON(http.StatusOK, response)
}
func DeleteAuthor(c *gin.Context) {
	var response model.Response
	db := config.Connect()

	authorID := c.Param("id")

	//เช็คว่ามีตัวที่ลบไหมจาก id
	if err := db.First(&model.Author{}, authorID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	// ลบ author
	if err := db.Table("author").Where("AUTHOR_ID = ?", authorID).Delete(&authorID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Status = http.StatusOK
	response.Message = "Author deleted successfully"

	c.JSON(http.StatusOK, response)
}
