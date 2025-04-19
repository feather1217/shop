package api

import (
	"backend/services"
	"encoding/json"
	"net/http"
	"net/url"
	"github.com/gin-gonic/gin"
)

func LineCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code"})
		return
	}

	user, err := services.HandleLineCallback(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 將 user 轉為 JSON 並放入 URL
	userJson, _ := json.Marshal(user)
	redirectURL := "http://localhost:5173/callback?user=" + url.QueryEscape(string(userJson))

	c.Redirect(http.StatusFound, redirectURL)
}
