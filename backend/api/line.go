package api

import (
	"backend/models"
	"net/http"
	"backend/services"
	"github.com/gin-gonic/gin"
)
//Line 登入回調
func LineCallbackHandler(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not found"})
		return
	}

	loginRes, err := services.HandleLineCallback(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":        loginRes.Profile,
		"accessToken": loginRes.AccessToken,
	})
}

func SendLineMessageHandler(c *gin.Context) {
	var req models.LineMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "格式錯誤"})
		return
	}

	err := services.SendLineMessage(req.UserID, req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "message sent"})
	
}
