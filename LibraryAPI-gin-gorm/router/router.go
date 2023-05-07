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

	router.GET("/getAuthor", controller.GetAuthors)
	router.POST("/insertAuthor", controller.CreateAuthor)
	router.PUT("/updateAuthor", controller.UpdateAuthor)
	router.DELETE("/deleteAuthor/:id", controller.DeleteAuthor)
	// router.GET("/authorbyid/", authorController.GetAuthorById)
	return router
}
