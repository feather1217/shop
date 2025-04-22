package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backend/models"
	"backend/database"
	"gorm.io/gorm"
)

// CreateProduct 新增商品
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 自動計算價格區間
	min, max := models.CalculatePriceRange(product.Variants)
	product.PriceMin = min
	product.PriceMax = max

	// 儲存產品資料並預加載關聯的 SpecTypes 和 Variants
	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "新增失敗"})
		return
	}

	// 使用 Preload 加載關聯資料
	if err := database.DB.Preload("SpecTypes.Values").Preload("Variants").First(&product, product.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查詢關聯資料失敗"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// 獲取所有商品
func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := database.DB.Preload("SpecTypes.Values").Preload("Variants").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查詢失敗"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// 獲取單個商品
func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := database.DB.Preload("SpecTypes.Values").Preload("Variants").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "商品未找到"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// 更新商品
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 取得現有商品（確認是否存在）
	var existingProduct models.Product
	if err := database.DB.Preload("SpecTypes.Values").Preload("Variants").First(&existingProduct, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "商品未找到"})
		return
	}

	// 刪除舊的關聯資料（Variants 與 SpecTypes -> SpecValues）
	database.DB.Where("product_id = ?", id).Delete(&models.ProductVariant{})
	for _, specType := range existingProduct.SpecTypes {
		database.DB.Where("spec_type_id = ?", specType.ID).Delete(&models.SpecValue{})
	}
	database.DB.Where("product_id = ?", id).Delete(&models.SpecType{})

	// 更新主要欄位（商品名稱）
	existingProduct.Name = newProduct.Name

	// 更新關聯資料（Variants、SpecTypes、SpecValues）
	existingProduct.SpecTypes = newProduct.SpecTypes
	existingProduct.Variants = newProduct.Variants

	// 重新計算價格範圍
	min, max := models.CalculatePriceRange(newProduct.Variants)
	existingProduct.PriceMin = min
	existingProduct.PriceMax = max

	// 儲存（包含關聯）
	if err := database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&existingProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失敗"})
		return
	}

	c.JSON(http.StatusOK, existingProduct)
}

//刪除商品
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	// 查詢商品是否存在
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "商品未找到"})
		return
	}

	// 刪除商品
	if err := database.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "刪除失敗"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "商品已刪除"})
}

