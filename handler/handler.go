package handler

import (
	"errors"
	"keeper/keeper"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	FILE_SIGNATURE = "image"
)

func RegisterFileHandler(c *gin.Context) {
	multiPartForm, err := c.MultipartForm()
	if err != nil {
		handleError(c, err)
		return
	}
	arquivos := multiPartForm.File[FILE_SIGNATURE]

	if len(arquivos) != 1 {
		handleError(c, errors.New("more than one file sent"))
		return
	}
	if err = keeper.RegisterFile(arquivos[0]); err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"filename": arquivos[0].Filename})
}

func RemoveFileHandler(c *gin.Context) {
	fileName := c.Param("file_name")
	if err := keeper.RemoveFile(&fileName); err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func handleError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
	ctx.Abort()
}
