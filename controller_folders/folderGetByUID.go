package controller_folders

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/AndrewSalko/salkodev.edms.go/database_folders"
	"github.com/gin-gonic/gin"
)

func GetFolderByUID(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	uid := c.Param(controller.UIDParam)
	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "folder uid not specified"})
		return
	}

	folder, err := database_folders.FindFolderByUID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, folder)
}
