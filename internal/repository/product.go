package repository

import (
	"database/sql"
	"grocery-api/internal/model"
	"time"
)

func NewProductRepository(db *sql.DB) Product {
	return &Handler{db}
}

func (h Handler) GetCategories() ([]model.ProductCategories, error) {
	var categories []model.ProductCategories
	rows, err := h.db.Query(getListCategories)
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var category model.ProductCategories
		err = rows.Scan(&category.ID, &category.CategoryName)
		if err != nil {
			return categories, err
		}
		categories = append(categories, category)
	}
	return categories, err
}

func (h Handler) GetProductsByCategory(id int) ([]model.Product, error) {
	var products []model.Product
	rows, err := h.db.Query(getProductByCategory, id)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		err = rows.Scan(&product.ID, &product.ProductName, &product.Category, &product.Qty, &product.PromoPrice, &product.SellingPrice)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, err
}

func (h Handler) GetProduct() ([]model.Product, error) {
	var products []model.Product
	rows, err := h.db.Query(getProducts)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		err = rows.Scan(&product.ID, &product.ProductName, &product.Category, &product.Qty, &product.PromoPrice, &product.SellingPrice)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, err
}

func (h Handler) GetProductDetail(id int) (model.Product, error) {
	var product model.Product
	rows, err := h.db.Query(getProductDetail, id)
	if err != nil {
		return product, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&product.ID, &product.ProductName, &product.Category, &product.Description, &product.Qty, &product.PromoPrice, &product.SellingPrice, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return product, err
		}
	}
	return product, err
}

func (h Handler) UpdateQuantityTx(tx *sql.Tx, product model.Product) error {
	_, err := tx.Exec(UpdateProductQuantity, product.Qty, time.Now(), product.ID)
	if err != nil {
		return err
	}
	return err
}

func (h Handler) AddProduct(product model.ProductParam) error {
	_, err := h.db.Exec(insertNewProduct, product.ProductName, product.CategoryID, product.Description, product.Qty, product.SellingPrice, product.PromoPrice)
	if err != nil {
		return err
	}
	return err
}
func (h Handler) UpdateProduct(product model.ProductParam) error {
	_, err := h.db.Exec(updateProduct, product.ProductName, product.CategoryID, product.Description, product.Qty, product.SellingPrice, product.PromoPrice, time.Now(), product.ID)
	if err != nil {
		return err
	}
	return err
}
