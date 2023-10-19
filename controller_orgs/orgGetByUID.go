package controller_orgs

import (
	"context"
	"net/http"
	"time"

	"github.com/AndrewSalko/salkodev.edms.go/controller"
	"github.com/AndrewSalko/salkodev.edms.go/database_orgs"
	"github.com/gin-gonic/gin"
)

func GetOrganizationByUID(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	uid := c.Param(controller.UIDParam)
	if uid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "organization uid not specified"})
		return
	}

	org, err := database_orgs.FindOrganizationByUID(ctx, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, org)
}
