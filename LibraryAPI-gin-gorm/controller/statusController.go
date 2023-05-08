package controller

import (
	"LibraryAPI/config"
	"LibraryAPI/model"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatus(c *gin.Context) {
	var status []model.Status
	var response model.StatusResponse
	var responseError model.Response

	statusId := c.Query("statusId")
	statusName := c.Query("statusName")
	statusCode := c.Query("statusCode")

	db := config.Connect()
	if err := db.Table("status").Where("(STATUS_NAME LIKE ? OR ? = '' ) AND (STATUS_ID = ? OR ? = '') AND (STATUS_CODE = ? OR ? = '') ", "%"+statusName+"%", statusName, statusId, statusId, statusCode, statusCode).Order("STATUS_ID asc").Find(&status).Error; err != nil {
		responseError.Status = http.StatusInternalServerError
		responseError.Message = err.Error()
		c.JSON(http.StatusInternalServerError, responseError)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = status
	c.JSON(http.StatusOK, response)
}

func InsertStatus(c *gin.Context) {
	var status model.Status
	var response model.Response

	if err := c.BindJSON(&status); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	if status.STATUS_CODE == nil || status.STATUS_NAME == nil {
		response.Status = http.StatusNotFound
		response.Message = "Incomplete input values"
		c.JSON(http.StatusNotFound, response)
		return
	}
	db := config.Connect()

	if err := db.Table("status").Create(&status).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Insert Status Success"
	c.JSON(http.StatusOK, response)
}

func UpdateStatus(c *gin.Context) {
	var status model.Status
	var response model.Response

	if err := c.BindJSON(&status); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	db := config.Connect()
	if err := db.Table("status").Where("STATUS_ID = ?", status.STATUS_ID).First(&model.Status{}).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "Status not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	if err := db.Table("status").Where("STATUS_ID = ?", status.STATUS_ID).Updates(map[string]interface{}{"STATUS_NAME": status.STATUS_NAME, "STATUS_CODE": status.STATUS_CODE}).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Status updated successfully"

	c.JSON(http.StatusOK, response)
}

func DeleteStatus(c *gin.Context) {
	var status model.Status
	var response model.Response
	statusId := c.Query("statusId")
	db := config.Connect()
	if err := db.Table("status").Where("STATUS_ID = ?", statusId).First(&status).Error; err != nil {
		response.Status = http.StatusNotFound
		response.Message = "Status not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	if err := db.Table("status").Where("STATUS_ID = ?", statusId).Delete(&statusId).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Status Delete successfully"

	c.JSON(http.StatusOK, response)
}
