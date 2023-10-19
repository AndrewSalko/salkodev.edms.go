package controller_users

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/AndrewSalko/salkodev.edms.go/database_users"
	"github.com/gin-gonic/gin"
)

func GetUserByUID(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	uid := c.Param(controller.UIDParam)
	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user uid not specified"})
		return
	}

	user, err := database_users.FindUserByUID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
