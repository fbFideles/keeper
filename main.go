package main

import (
	"keeper/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	router(v1)

	r.Run(":8183")
}

func router(r *gin.RouterGroup) {
	r.DELETE("file/:file_name", handler.RemoveFileHandler)
	r.POST("files", handler.RegisterFileHandler)
}
