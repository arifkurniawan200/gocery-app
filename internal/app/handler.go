package app

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"grocery-api/internal"
	"grocery-api/internal/model"
	"net/http"
	"strconv"
	"time"
)

type jwtCustomClaims struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	ID       int64  `json:"id"`
	jwt.RegisteredClaims
}

type ResponseSuccess struct {
	Messages string      `json:"messages,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

type ResponseFailed struct {
	Messages string      `json:"messages,omitempty"`
	Error    interface{} `json:"error,omitempty"`
}

func (u Handler) RegisterUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	hash, err := internal.HashPassword(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}
	err = u.User.RegisterUser(model.User{
		Username:     username,
		PasswordHash: hash,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to register user",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, ResponseSuccess{
		Messages: "success register user",
	})
}

func (u Handler) LoginUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	userInfo, err := u.User.VerifyUser(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to login",
			Error:    err.Error(),
		})
	}
	if !internal.VerifyPassword(password, userInfo.PasswordHash) {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "invalid username/password",
			Error:    "username or password is mismatch",
		})
	}
	claims := &jwtCustomClaims{
		userInfo.Username,
		userInfo.IsAdmin,
		userInfo.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "error when generate token",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   accessToken,
	})
}

func (u Handler) GetListProduct(c echo.Context) error {
	var (
		data []model.Product
		err  error
	)

	idParam := c.QueryParam("id")

	if idParam != "" {
		id, _ := strconv.Atoi(idParam)
		data, err = u.Product.GetProductsByCategory(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ResponseFailed{
				Messages: "error when get categories",
				Error:    err.Error(),
			})
		}
	} else {
		data, err = u.Product.GetProduct()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ResponseFailed{
				Messages: "error when get categories",
				Error:    err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch products",
		Data:     data,
	})
}

func (u Handler) GetListCategories(c echo.Context) error {
	data, err := u.Product.GetCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "error when get categories",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch categories",
		Data:     data,
	})
}

func (u Handler) GetProductDetail(c echo.Context) error {
	idParam := c.QueryParam("id")
	id, _ := strconv.Atoi(idParam)
	data, err := u.Product.GetProductDetail(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "error when get categories",
			Error:    err.Error(),
		})
	}
	if data.ID == 0 {
		return c.JSON(http.StatusNotFound, ResponseFailed{
			Messages: "data not found",
			Error:    nil,
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch detail product",
		Data:     data,
	})
}

func (u Handler) AddToCart(c echo.Context) error {
	var cart model.AddToCartParam

	if err := c.Bind(&cart); err != nil {
		return c.JSON(http.StatusBadRequest, ResponseFailed{
			Messages: "Failed to parse request body",
			Error:    err.Error(),
		})
	}

	// get user id from token
	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid or missing claims",
		})
	}

	// Check if the "is_admin" claim exists and is true
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get user id",
		})
	}

	// assign user id from context to struct
	cart.UserID = int64(userID)

	detail, err := u.Product.GetProductDetail(cart.ProductID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "Failed to get database",
			Error:    err.Error(),
		})
	}

	// check product stoct is available
	if cart.Quantity > detail.Qty {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "total product is not enough",
			Error:    "current total product is not enough add to cart",
		})
	}

	// start process using transaction
	tx, err := u.User.BeginTx()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "Failed to get database",
			Error:    err.Error(),
		})
	}

	err = u.User.AddUserCartTx(tx, cart)
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "Failed to get database",
			Error:    err.Error(),
		})
	}

	err = u.Product.UpdateQuantityTx(tx, model.Product{
		ID:  detail.ID,
		Qty: detail.Qty - cart.Quantity,
	})
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "Failed to get database",
			Error:    err.Error(),
		})
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "Failed to get database",
			Error:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success add to cart",
	})
}

func (u Handler) GetUserCarts(c echo.Context) error {
	// get user id from token
	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid or missing claims",
		})
	}

	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "failed to get user id",
		})
	}
	data, err := u.User.UserCartList(int(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ResponseFailed{
			Messages: "error when get user carts",
			Error:    err.Error(),
		})
	}
	return c.JSON(http.StatusOK, ResponseSuccess{
		Messages: "success fetch users cart",
		Data:     data,
	})
}
