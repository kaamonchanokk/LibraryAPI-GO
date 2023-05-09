package controller

import (
	"LibraryAPI/config"
	"LibraryAPI/model"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context) {
	var category []model.Category
	var response model.CategoryResponse
	var responseError model.Response

	categoryId := c.Query("categoryId")
	categoryCode := c.Query("categoryCode")
	categoryName := c.Query("categoryName")

	db := config.Connect()

	if err := db.Table("category").Where(" (CATEGORY_ID = ? OR ? = '' ) AND (CATEGORY_CODE = ? OR ? = '' ) AND (CATEGORY_NAME LIKE ? OR ? = '') ", categoryId, categoryId, categoryCode, categoryCode, "%"+categoryName+"%", categoryName).Find(&category).Error; err != nil {
		responseError.Status = http.StatusInternalServerError
		responseError.Message = err.Error()
		c.JSON(http.StatusInternalServerError, responseError)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = category
	c.JSON(http.StatusOK, response)
}
func InsertCategory(c *gin.Context) {
	var category model.Category
	var response model.Response

	if err := c.BindJSON(&category); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	if category.CATEGORY_CODE == nil || category.CATEGORY_NAME == nil {
		response.Status = http.StatusNotFound
		response.Message = "Incomplete input values"
		c.JSON(http.StatusNotFound, response)
		return
	}
	db := config.Connect()

	if err := db.Table("category").Create(&category).Error; err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Success"
	c.JSON(http.StatusOK, response)

}
func UpdateCategory(c *gin.Context) {
	var category model.Category
	var response model.Response

	if err := c.BindJSON(&category); err != nil {
		response.Status = http.StatusBadRequest
		response.Message = err.Error()
		c.JSON(http.StatusNotFound, response)
		return
	}
	db := config.Connect()
	if err := db.Table("category").Where("CATEGORY_ID = ?",category.CATEGORY_ID).First(&model.Category{}).Error;err != nil {
		response.Status = http.StatusNotFound
		response.Message = "Category not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	if err := db.Table("category").Where("CATEGORY_ID = ?",category.CATEGORY_ID).Updates(map[string]interface{}{"CATEGORY_CODE":category.CATEGORY_CODE,"CATEGORY_NAME":category.CATEGORY_NAME}).Error ; err != nil{
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Category updated successfully"

	c.JSON(http.StatusOK, response)
}
func DeleteCategory(c *gin.Context) {
	var response model.Response

	categoryId := c.Query("categoryId")

	db := config.Connect()
	if err := db.Table("category").Where("CATEGORY_ID = ?",categoryId).First(&model.Category{}).Error;err != nil {
		response.Status = http.StatusNotFound
		response.Message = "Category not found"
		c.JSON(http.StatusNotFound, response)
		return
	}
	if err := db.Table("category").Where("CATEGORY_ID = ?",categoryId).Delete(&model.Category{}).Error;err != nil {
		response.Status = http.StatusInternalServerError
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}


	response.Status = http.StatusOK
	response.Message = "Category Delete successfully"

	c.JSON(http.StatusOK, response)
}
