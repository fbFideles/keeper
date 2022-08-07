package main

import (
	"keeper/handler"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	router(v1)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8183"
	}

	r.Run(":"+port)
}

func router(r *gin.RouterGroup) {
	r.DELETE("file/:file_name", handler.RemoveFileHandler)
	r.POST("files", handler.RegisterFileHandler)
}
