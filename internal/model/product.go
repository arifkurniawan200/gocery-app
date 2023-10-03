package model

import "time"

type Product struct {
	ID           int64      `json:"id,omitempty" db:"id"`
	ProductName  string     `json:"product_name,omitempty" db:"product_name"`
	Category     string     `json:"category,omitempty" db:"category"`
	Description  string     `json:"description,omitempty" db:"description"`
	Qty          int        `json:"qty,omitempty" db:"qty"`
	SellingPrice float64    `json:"selling_price,omitempty" db:"selling_price"`
	PromoPrice   float64    `json:"promo_price,omitempty" db:"promo_price"`
	CreatedAt    *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type ProductCategories struct {
	ID           int64  `json:"id,omitempty" db:"id"`
	CategoryName string `json:"category_name,omitempty" db:"category_name"`
}

type AddToCartParam struct {
	UserID    int64
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ProductParam struct {
	ID           int64
	ProductName  string  `json:"product_name,omitempty" db:"product_name"`
	CategoryID   int     `json:"category_id" db:"category_id"`
	Description  string  `json:"description,omitempty" db:"description"`
	Qty          int     `json:"qty,omitempty" db:"qty"`
	SellingPrice float64 `json:"selling_price,omitempty" db:"selling_price"`
	PromoPrice   float64 `json:"promo_price,omitempty" db:"promo_price"`
}
