package controller

import (
	"LibraryAPI/config"
	"LibraryAPI/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var getBook []model.GetBook
	var response model.BookResponse
	var responseError model.Response
	authorCode := c.Query("authorCode")
	authorName := c.Query("authorName")
	bookName := c.Query("bookName")
	bookIsbn := c.Query("bookIsbn")
	bookId := c.Query("bookId")
	categoryCode := c.Query("categoryCode")
	categoryName := c.Query("categoryName")
	db := config.Connect()
	if err := db.Table("book").Preload("Author").Preload("Category").Joins("INNER JOIN author ON book.AUTHOR_ID = author.AUTHOR_ID").
		Joins("INNER JOIN category ON book.CATEGORY_ID = category.CATEGORY_ID").Where("(author.AUTHOR_CODE = ? OR  ? = '') AND (author.AUTHOR_NAME LIKE ? OR  ? = '') AND (book.BOOK_ISBN = ? OR  ? = '') AND (book.BOOK_ISBN = ? OR  ? = '') AND (book.BOOK_NAME LIKE ? OR  ? = '') AND (category.CATEGORY_CODE = ? OR  ? = '') AND (category.CATEGORY_NAME LIKE ? OR  ? = '')", authorCode, authorCode, "%"+authorName+"%", authorName, bookId, bookId, bookIsbn, bookIsbn, "%"+bookName+"%", bookName, categoryCode, categoryCode, "%"+categoryName+"%", categoryName).Find(&getBook).Error; err != nil {
		responseError.Status = http.StatusInternalServerError
		responseError.Message = err.Error()
		c.JSON(http.StatusInternalServerError, responseError)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = getBook
	c.JSON(http.StatusOK, response)
}

func InsertBook(c *gin.Context) {
	var book model.Book
	var response model.Response

	if err := c.BindJSON(&book); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	if book.BOOK_ISBN == nil || book.BOOK_NAME == nil || book.BOOK_QUANILTY == nil || book.AUTHOR_CODE == nil || book.CATEGORY_CODE == nil  {
		response.Status = http.StatusNotFound
		response.Message = "Incomplete input values"
		c.JSON(http.StatusNotFound, response)
		return
	}
	// fmt.Println(book)
	db := config.Connect()

	//find AuthorId and CategoryId
	var authorId int
	var categoryId int
	if err := db.Table("author").Where("AUTHOR_CODE = ? ", book.AUTHOR_CODE).Select("author.AUTHOR_ID").Scan(&authorId).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if err := db.Table("category").Where("CATEGORY_CODE = ? ", book.CATEGORY_CODE).Select("category.CATEGORY_ID").Scan(&categoryId).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	fmt.Println(categoryId)
	book.AUTHOR_ID = &authorId
	book.CATEGORY_ID = &categoryId

	//find AuthorId and CategoryId
	if err := db.Table("book").Exec("INSERT INTO book (BOOK_ID, BOOK_ISBN, BOOK_NAME, BOOK_QUANILTY, AUTHOR_ID, CATEGORY_ID) VALUES (?, ?, ?, ?, ?, ?)",book.BOOK_ID,book.BOOK_ISBN,book.BOOK_NAME,book.BOOK_QUANILTY,book.AUTHOR_ID,book.CATEGORY_ID).Error;err!=nil{
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Insert Book Success"
	c.JSON(http.StatusOK, response)
}
func UpdateBook(c *gin.Context){
	var book model.Book
	var response model.Response

	if err := c.BindJSON(&book); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	
	db := config.Connect()
	if err := db.Table("book").Where("BOOK_ID = ?", book.BOOK_ID).First(&model.Book{}).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "Book not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	//find AuthorId and CategoryId
	var authorId int
	var categoryId int
	if err := db.Table("author").Where("AUTHOR_CODE = ? ", book.AUTHOR_CODE).Select("author.AUTHOR_ID").Scan(&authorId).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if err := db.Table("category").Where("CATEGORY_CODE = ? ", book.CATEGORY_CODE).Select("category.CATEGORY_ID").Scan(&categoryId).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	fmt.Println(categoryId)
	book.AUTHOR_ID = &authorId
	book.CATEGORY_ID = &categoryId

	if err := db.Table("book").Where("BOOK_ID = ?", book.BOOK_ID).Updates(map[string]interface{}{"BOOK_ISBN": book.BOOK_ISBN,"BOOK_NAME": book.BOOK_NAME, "BOOK_QUANILTY": book.BOOK_QUANILTY, "AUTHOR_ID": book.AUTHOR_ID, "CATEGORY_ID": book.CATEGORY_ID}).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Book updated successfully"

	c.JSON(http.StatusOK, response)

}
func DeleteBook(c *gin.Context){
	var response model.Response
	bookId := c.Query("bookId")
	
	db := config.Connect()
	if err := db.Table("book").Where("BOOK_ID = ?", bookId).First(&model.Book{}).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "Book not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	if err:=db.Table("book").Where("BOOK_ID = ?", bookId).Delete(&model.Book{}).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Book Delete successfully"

	c.JSON(http.StatusOK, response)
}