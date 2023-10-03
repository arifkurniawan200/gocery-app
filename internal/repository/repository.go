package repository

import (
	"database/sql"
	"grocery-api/internal/model"
)

type User interface {
	BeginTx() (*sql.Tx, error)
	RegisterUser(user model.User) error
	VerifyUser(username string) (model.User, error)
	AddUserCartTx(tx *sql.Tx, param model.AddToCartParam) error
	UserCartList(userID int) ([]model.Product, error)
}

type Product interface {
	GetCategories() ([]model.ProductCategories, error)
	GetProductsByCategory(id int) ([]model.Product, error)
	GetProduct() ([]model.Product, error)
	GetProductDetail(id int) (model.Product, error)
	UpdateQuantityTx(tx *sql.Tx, param model.Product) error
	AddProduct(product model.ProductParam) error
	UpdateProduct(product model.ProductParam) error
}
