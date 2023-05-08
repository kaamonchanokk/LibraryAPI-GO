package controller

import (
	"LibraryAPI/config"
	"LibraryAPI/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	var author []model.Author
	var response model.AuthorResponse
	var responseError model.Response
	authorId := c.Query("authorId")
	authorCode := c.Query("authorCode")
	authorName := c.Query("authorName")

	db := config.Connect()

	if err := db.Table("author").Where("(AUTHOR_NAME LIKE ? OR ? = '' ) AND (AUTHOR_ID = ? OR ? = '') AND (AUTHOR_CODE = ? OR ? = '') ", "%"+authorName+"%", authorName, authorId, authorId, authorCode, authorCode).Order("AUTHOR_ID asc").Find(&author).Error; err != nil {
		responseError.Status = http.StatusInternalServerError
		responseError.Message = err.Error()
		c.JSON(http.StatusInternalServerError, responseError)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = author
	c.JSON(http.StatusOK, response)
}

func CreateAuthor(c *gin.Context) {
	var author model.Author
	var response model.Response
	var count int64
	var authorCode string
	//จาก JSON -> author
	if err := c.BindJSON(&author); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	//เช็คค่าว่ามีครบไหม
	if author.AUTHOR_ADDRESS == nil || author.AUTHOR_NAME == nil {
		response.Status = http.StatusNotFound
		response.Message = "Incomplete input values"
		c.JSON(http.StatusNotFound, response)
		return
	}
	db := config.Connect()

	//gen authorCode
	if err := db.Table("author").Count(&count).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	authorCode = fmt.Sprintf("A%03d", count+1)
	author.AUTHOR_CODE = &authorCode

	//เพิ่มค่า
	if err := db.Table("author").Create(&author).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusCreated
	response.Message = "Insert Author Success"
	c.JSON(http.StatusCreated, response)
}

func UpdateAuthor(c *gin.Context) {
	var author model.Author
	var response model.Response
	db := config.Connect()

	//จาก JSON -> author
	if err := c.BindJSON(&author); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	// fmt.Println(author)
	//เช็คว่ามีตัวที่แก้ไขไหมจาก id
	if err := db.Table("author").First(&model.Author{}, author.AUTHOR_ID).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "Author not found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	//บันทึกการแก้ไข
	fmt.Println(author)
	// if err := db.Table("author").Save(&author).Error; err != nil {
	// 	response.Status = http.StatusInternalServerError
	// 	response.Message = err.Error()
	// 	c.JSON(http.StatusInternalServerError, response)
	// 	return
	// }
	if err := db.Table("author").Where("AUTHOR_ID = ?", author.AUTHOR_ID).Updates(map[string]interface{}{"AUTHOR_NAME": author.AUTHOR_NAME, "AUTHOR_ADDRESS": author.AUTHOR_ADDRESS}).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Author updated successfully"

	c.JSON(http.StatusOK, response)
}

func DeleteAuthor(c *gin.Context) {
	var response model.Response
	var author model.Author

	db := config.Connect()

	authorId := c.Query("authorId")

	//เช็คว่ามีตัวที่ลบไหมจาก Code
	if err := db.Where("AUTHOR_ID = ?", authorId).First(&author).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "Author not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	fmt.Println(author)
	// ลบ author
	if err := db.Table("author").Where("AUTHOR_ID = ?", authorId).Delete(&authorId).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Author deleted successfully"

	c.JSON(http.StatusOK, response)
}

//เวอร์ชั่นลบแบบ Id
// func DeleteAuthor(c *gin.Context) {
// 	var response model.Response
// 	db := config.Connect()

// 	authorID := c.Param("id")

// 	//เช็คว่ามีตัวที่ลบไหมจาก id
// 	if err := db.First(&model.Author{}, authorID).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
// 		return
// 	}

// 	// ลบ author
// 	if err := db.Table("author").Where("AUTHOR_ID = ?", authorID).Delete(&authorID).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	response.Status = http.StatusOK
// 	response.Message = "Author deleted successfully"

// 	c.JSON(http.StatusOK, response)
// }
