// models/product.go
package models

import "gorm.io/gorm"

type Product struct {
	ID         uint           `gorm:"primaryKey"`
	Name       string         `json:"name"`
	PriceMin   float64        `json:"priceMin"`
	PriceMax   float64        `json:"priceMax"`
	SpecTypes  []SpecType     `gorm:"foreignKey:ProductID" json:"specTypes"`
	Variants   []ProductVariant `gorm:"foreignKey:ProductID" json:"variants"`
	CreatedAt  int64
	UpdatedAt  int64
}

type SpecType struct {
	gorm.Model
	ProductID uint         `json:"-"`
	Name      string       `json:"name"`
	Values    []SpecValue  `gorm:"foreignKey:SpecTypeID" json:"values"`
}

type SpecValue struct {
	gorm.Model
	SpecTypeID uint   `json:"-"`
	Value      string `json:"value"`
	ImageURL   string `json:"imageUrl,omitempty"` // 只有第一種規格會填
}

type ProductVariant struct {
	gorm.Model
	ProductID   uint    `json:"-"`
	SpecValue1  string  `json:"specValue1"`
	SpecValue2  string  `json:"specValue2,omitempty"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

func CalculatePriceRange(variants []ProductVariant) (min, max float64) {
	if len(variants) == 0 {
		return 0, 0
	}
	min, max = variants[0].Price, variants[0].Price
	for _, v := range variants {
		if v.Price < min {
			min = v.Price
		}
		if v.Price > max {
			max = v.Price
		}
	}
	return
}