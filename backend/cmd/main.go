package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "backend/api"
    "time"
)

func main() {
    r := gin.Default()

    // CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))
    r.GET("/api/line/callback", api.LineCallbackHandler)
    r.POST("/api/line/message", api.SendLineMessageHandler)

	// Routes
    r.Run(":8080")
}
