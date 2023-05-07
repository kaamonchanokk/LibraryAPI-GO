package main

import (
    // "github.com/gin-gonic/gin"
    "LibraryAPI/router"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := router.Router()
    r.Run(":8080")
}

