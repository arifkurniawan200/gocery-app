package repository

import (
	"context"
	"database/sql"
	"grocery-api/internal/model"
)

type Handler struct {
	db *sql.DB
}

func (h Handler) BeginTx() (*sql.Tx, error) {
	return h.db.BeginTx(context.Background(), &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
}

func NewUserRepository(db *sql.DB) User {
	return &Handler{db}
}

func (h Handler) RegisterUser(user model.User) error {
	_, err := h.db.Exec(registerUser, user.Username, user.PasswordHash, user.IsAdmin)
	if err != nil {
		return err
	}
	return nil
}

func (h Handler) AddUserCartTx(tx *sql.Tx, param model.AddToCartParam) error {
	_, err := tx.Exec(addToCart, param.UserID, param.ProductID, param.Quantity)
	if err != nil {
		return err
	}
	return err
}

func (h Handler) UserCartList(userID int) ([]model.Product, error) {
	var products []model.Product
	rows, err := h.db.Query(getUserCart, userID)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		err = rows.Scan(&product.ID, &product.ProductName, &product.Qty, &product.PromoPrice, &product.SellingPrice, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, err

}

func (h Handler) VerifyUser(username string) (model.User, error) {
	var user model.User
	rows, err := h.db.Query(getUserByName, username)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.IsAdmin)
		if err != nil {
			return user, err
		}
	}
	return user, err
}
