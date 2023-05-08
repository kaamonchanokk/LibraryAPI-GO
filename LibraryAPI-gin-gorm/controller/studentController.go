package controller

import (
	"LibraryAPI/config"
	"LibraryAPI/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetStudent(c *gin.Context) {
	var student []model.Student
	var response model.StudentResponse
	var responseError model.Response

	studentId := c.Query("studentId")
	studentName := c.Query("studentName")
	studentCode := c.Query("studentCode")
	studentYear := c.Query("studentYear")

	db := config.Connect()
	if err := db.Table("student").Where("(STUDENT_NAME LIKE ? OR ? = '' ) AND (STUDENT_ID = ? OR ? = '') AND (STUDENT_CODE = ? OR ? = '') AND (STUDENT_YEAR = ? OR ? = '') ", "%"+studentName+"%", studentName, studentId, studentId, studentCode, studentCode, studentYear, studentYear).Order("STUDENT_ID asc").Find(&student).Error; err != nil {
		responseError.Status = http.StatusInternalServerError
		responseError.Message = err.Error()
		c.JSON(http.StatusInternalServerError, responseError)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = student
	c.JSON(http.StatusOK, response)
}
func InsertStudent(c *gin.Context) {
	var student model.Student
	var response model.Response
	var count int64
	var studentCode string
	now := time.Now()
	thaiYear := now.Year() + 543

	if err := c.BindJSON(&student); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	if student.STUDENT_NAME == nil {
		response.Status = http.StatusNotFound
		response.Message = "Incomplete input values"
		c.JSON(http.StatusNotFound, response)
		return
	}
	db := config.Connect()

	// gen studentCode
	if err := db.Table("student").Where("STUDENT_YEAR = ? ", thaiYear).Count(&count).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	studentCode = fmt.Sprintf("%d%08d", thaiYear%100, count+1)
	student.STUDENT_CODE = &studentCode
	student.STUDENT_YEAR = &thaiYear

	if err := db.Table("student").Create(&student).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Insert Student Success"
	c.JSON(http.StatusOK, response)
}

func UpdateStudent(c *gin.Context) {
	var student model.Student
	var response model.Response
	if err := c.BindJSON(&student); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	db := config.Connect()

	if err := db.Table("student").First(&model.Student{}, student.STUDENT_ID).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "Student not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	if err := db.Table("student").Where("STUDENT_ID = ? ", student.STUDENT_ID).Update("STUDENT_NAME", student.STUDENT_NAME).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Student updated successfully"
	c.JSON(http.StatusOK, response)
}
func DeleteStudent(c *gin.Context) {
	var response model.Response
	var student model.Student
	db := config.Connect()
	studentId := c.Query("studentId")

	if err := db.Table("student").Where("STUDENT_ID = ? ", studentId).First(&student).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "Student not found"
		c.JSON(http.StatusNotFound, response)
		return
	}

	if err := db.Table("student").Where("STUDENT_ID = ? ", studentId).Delete(&studentId).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Message = "Student deleted successfully"

	c.JSON(http.StatusOK, response)
}
