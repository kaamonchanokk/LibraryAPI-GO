package controller

import (
	"LibraryAPI/config"
	"LibraryAPI/model"
	// "fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBorrowReport(c *gin.Context) {
	var getBorrow []model.GetBorrow
	var response model.BorrowResponse
	var responseError model.Response

	borrowId := c.Query("borrowId")
	bookName := c.Query("bookName")
	studentName := c.Query("studentName")
	statusCode := c.Query("statusCode")

	db := config.Connect()
	if err := db.Table("borrow").Joins("INNER JOIN book ON book.BOOK_ID = borrow.BOOK_ID").Joins("INNER JOIN student ON student.STUDENT_ID = borrow.STUDENT_ID").Joins("INNER JOIN status ON status.STATUS_ID = borrow.STATUS_ID").Select("borrow.BORROW_ID,book.BOOK_NAME,student.STUDENT_NAME,borrow.DATE_BORROW,borrow.DATE_RETURN,borrow.BORROW_QUANTITY,status.STATUS_NAME").Where("(borrow.BORROW_ID = ? OR  ? = '') AND (book.BOOK_NAME LIKE ? OR  ? = '') AND (student.STUDENT_NAME LIKE ? OR  ? = '') AND (status.STATUS_CODE = ? OR  ? = '') ", borrowId, borrowId, "%"+bookName+"%", bookName, "%"+studentName+"%", studentName, statusCode, statusCode).Find(&getBorrow).Error; err != nil {
		responseError.Status = http.StatusInternalServerError
		responseError.Message = err.Error()
		c.JSON(http.StatusInternalServerError, responseError)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = getBorrow
	c.JSON(http.StatusOK, response)
}
func InsertBorrow(c *gin.Context){
	var borrow model.Borrow
	var response model.Response
	var bookId int
	var studentId int
	var TotalBorrow int
	if err := c.BindJSON(&borrow); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	db := config.Connect()
	if err := db.Table("book").Where("BOOK_ISBN = ?",borrow.BOOK_ISBN).Select("BOOK_ID").Scan(&bookId).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "BookId not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	if err := db.Table("student").Where("STUDENT_CODE = ?",borrow.STUDENT_CODE).Select("STUDENT_ID").Scan(&studentId).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "StudentId not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	borrow.BOOK_ID = bookId
	borrow.STUDENT_ID = studentId
	if err := db.Table("book").Joins("LEFT JOIN ( SELECT br.BOOK_ID,SUM(br.BORROW_QUANTITY) AS BORROW_QUANTITY FROM borrow br WHERE br.STATUS_ID = '1' GROUP BY br.BOOK_ID) AS br ON book.BOOK_ID = br.BOOK_ID"). Select("IFNULL(BOOK_QUANILTY - BORROW_QUANTITY,0) as Total").Where("book.BOOK_ID = ?",bookId).Scan(&TotalBorrow).Error ; err != nil{
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if TotalBorrow == 0 || TotalBorrow < borrow.BORROW_QUANTITY {
		response.Status = http.StatusInternalServerError
		response.Message = "Not enough of Book"
		c.JSON(http.StatusNotFound, response)
		return
	}
	if borrow.DATE_BORROW.After(borrow.DATE_RETURN.Time) {
		// วันที่ DateReturn มากกว่า DateBorrow
		response.Status = http.StatusInternalServerError
		response.Message = "DATE_BORROW should before DATE_RETURN "
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if err := db.Table("borrow").Exec("INSERT INTO borrow (BORROW_ID , BOOK_ID, STUDENT_ID , DATE_BORROW, DATE_RETURN, BORROW_QUANTITY, STATUS_ID) VALUES (? , ?, ? , ?, ?, ?, 1)",borrow.BORROW_ID,borrow.BOOK_ID,borrow.STUDENT_ID,borrow.DATE_BORROW.Time,borrow.DATE_RETURN.Time,borrow.BORROW_QUANTITY).Error ; err != nil{
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Insert Borrow Success"
	c.JSON(http.StatusOK, response)
	
}
func UpdateDateBorrow(c *gin.Context){
	var borrow model.Borrow
	var response model.Response
	if err := c.BindJSON(&borrow); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	db := config.Connect()
	if err := db.Table("borrow").Where("BORROW_ID = ?",borrow.BORROW_ID).Select("BORROW_ID").First(&model.Borrow{}).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "BorrowId not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	if borrow.DATE_BORROW.After(borrow.DATE_RETURN.Time) {
		// วันที่ DateReturn มากกว่า DateBorrow
		response.Status = http.StatusInternalServerError
		response.Message = "DATE_BORROW should before DATE_RETURN "
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if err := db.Table("borrow").Where("BORROW_ID = ?",borrow.BORROW_ID).Updates(map[string]interface{}{"DATE_BORROW": borrow.DATE_BORROW.Time,"DATE_RETURN": borrow.DATE_RETURN.Time}).Error; err != nil{
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Update Borrow Success"
	c.JSON(http.StatusOK, response)
}
func UpdateStatusBorrow(c *gin.Context){
	var borrow model.Borrow
	var response model.Response
	var statusId int
	if err := c.BindJSON(&borrow); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	db := config.Connect()
	if err := db.Table("borrow").Where("BORROW_ID = ?",borrow.BORROW_ID).Select("BORROW_ID").First(&model.Borrow{}).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "BorrowId not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	if err := db.Table("status").Where("STATUS_CODE = ?",borrow.STATUS_CODE).Select("STATUS_ID").Scan(&statusId).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "StatusCode not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	borrow.STATUS_ID = statusId
	if err := db.Table("borrow").Where("BORROW_ID = ?",borrow.BORROW_ID).Update("STATUS_ID", borrow.STATUS_ID).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Update Borrow Status Success"
	c.JSON(http.StatusOK, response)
}
func DeleteBorrow(c *gin.Context) {
	var response model.Response
	borrowId := c.Query("borrowId")
	db := config.Connect()
	if err := db.Table("borrow").Where("BORROW_ID = ? ", borrowId).First(&model.Borrow{}).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "BorrowId not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	if err := db.Table("borrow").Where("BORROW_ID = ?",borrowId).Delete(&model.Borrow{}).Error ; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Delete Borrow Success"
	c.JSON(http.StatusOK, response)
}
