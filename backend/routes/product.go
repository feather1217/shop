// routes/product.go
package routes

import (
	"github.com/gin-gonic/gin"
	"backend/api"
)

func RegisterProductRoutes(r *gin.Engine) {
	product := r.Group("/api/products")
	{
		product.POST("/", api.CreateProduct)
		product.GET("/", api.GetProducts) // 獲取所有商品
		product.GET("/:id", api.GetProductByID) // 獲取單個商品
		product.PUT("/:id", api.UpdateProduct)   // 更新商品
		product.DELETE("/:id", api.DeleteProduct) // 刪除商品
	}
}
