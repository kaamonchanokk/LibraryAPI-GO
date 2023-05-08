package router

import (
	"LibraryAPI/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies([]string{"127.0.0.1/8"})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//----------------Author------------------------//

	router.GET("/getAuthor", controller.GetAuthors)
	router.POST("/insertAuthor", controller.CreateAuthor)
	router.PUT("/updateAuthor", controller.UpdateAuthor)
	router.DELETE("/deleteAuthor", controller.DeleteAuthor)
	// router.GET("/authorbyid/", authorController.GetAuthorById)

	//----------------Student------------------------//
	router.GET("/getStudent", controller.GetStudent)
	router.POST("/insertStudent", controller.InsertStudent)
	router.PUT("/updateStudent", controller.UpdateStudent)
	router.DELETE("/deleteStudent",controller.DeleteStudent)

	//----------------Status------------------------//
	router.GET("/getStatus", controller.GetStatus)
	router.POST("/insertStatus", controller.InsertStatus)
	router.PUT("/updateStatus", controller.UpdateStatus)
	router.DELETE("/deleteStatus",controller.DeleteStatus)

	return router
}
