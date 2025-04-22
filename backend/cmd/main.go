package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "backend/database"
    "backend/routes"
    "time"
)

func main() {
    dsn := "user=postgres password=12521252 host=localhost port=5432 dbname=postgres sslmode=disable" //資料庫
	database.InitDB(dsn) //去連接資料庫

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

    lineRoutes := r.Group("/api/line") // Line API 路由
    routes.LineRoutes(lineRoutes)
    routes.ProductRoutes(r)
	// Routes
    r.Run(":8080")
}
