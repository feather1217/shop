package routes

import (
	"github.com/gin-gonic/gin"
	"backend/api"
)

// LineRoutes 用來定義與 Line 相關的路由
func LineRoutes(r *gin.RouterGroup) {
	r.GET("/callback", api.LineCallbackHandler)
	r.POST("/message", api.SendLineMessageHandler)
}
