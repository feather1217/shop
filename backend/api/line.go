package api

import (

	"net/http"
	"backend/services"
	"github.com/gin-gonic/gin"
)

func LineCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not found"})
		return
	}

	profile, err := services.HandleLineCallback(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": profile,
	})
}

