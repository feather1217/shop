package routes

import (
	"github.com/gin-gonic/gin"
	"backend/api"
)

func LineRoutes(r *gin.RouterGroup) {
		r.GET("/callback", api.LineCallbackHandler)
		r.POST("/message", api.SendLineMessageHandler)
}