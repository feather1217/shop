package api

import (
	"backend/models"
	"net/http"
	"backend/services"
	"github.com/gin-gonic/gin"
)
//Line 登入回調
func LineCallbackHandler(c *gin.Context) {
	//接收前端傳過來的code
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code not found"}) //返回錯誤400
		return
	}
	//把code傳給lineServices處理
	loginRes, err := services.HandleLineCallback(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //返回錯誤500
		return
	}
	//成功
	c.JSON(http.StatusOK, gin.H{
		"user":        loginRes.Profile,
		"accessToken": loginRes.AccessToken,
	})
}

//發送Line消息
func SendLineMessageHandler(c *gin.Context) {
	//接收前端傳過來的請求
	var req models.LineMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "格式錯誤"}) //返回錯誤400
		return
	}
	//把請求傳給lineServices處理
	err := services.SendLineMessage(req.UserID, req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) //返回錯誤500
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "message sent"})
	
}
