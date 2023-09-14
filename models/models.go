package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ShopID            int    `gorm:"null" form:"shop_id" json:"shop_id"`
	Name              string `gorm:"null" form:"name" json:"name"`
	Description       string `gorm:"null" form:"description" json:"description"`
	ThumbnailUrl      string `gorm:"not null" form:"thumbnail_url" json:"thumbnail_url"`
	OriginPrice       int    `gorm:"not null" form:"origin_price" json:"origin_price"`
	DiscountedPrice   int    `gorm:"not null" form:"discounted_price" json:"discounted_price"`
	DiscountedRate    int    `gorm:"null" form:"discounted_rate" json:"discounted_rate"`
	Status            string `gorm:"not null" form:"status" json:"status"`
	InStock           uint64 `gorm:"not null" form:"in_stock" json:"in_stock"`
	IsPreorder        uint64 `gorm:"not null" form:"is_preorder" json:"is_preorder"`
	IsPurchasable     uint64 `gorm:"not null" form:"is_purchasable" json:"is_purchasable"`
	DeliveryCondition string `gorm:"null" form:"delivery_condition" json:"delivery_condition"`
	DeliveryDisplay   string `gorm:"null" form:"delivery_display" json:"delivery_display"`
}

type Users struct {
	gorm.Model
	Name     string `gorm:"null" form:"name" json:"name"`
	Email    string `gorm:"null" form:"email" json:"email"`
	Password string `gorm:"null" form:"password" json:"password"`
	Phone    string `gorm:"null" form:"phone" json:"phone"`
	Status   string `gorm:"not null" form:"status" json:"status"`
}

type Favorites struct {
	gorm.Model
	Code   string  `gorm:"not null" form:"code" json:"code"` //username (one user will have one cart code = username)
	ProdID int     `gorm:"not null" form:"prod_id" json:"prod_id"`
	Name   string  `gorm:"not null" form:"name" json:"name"`
	Price  float64 `gorm:"not null" form:"price" json:"price"`
	Items  int     `gorm:"not null" form:"items" json:"items"`
}
