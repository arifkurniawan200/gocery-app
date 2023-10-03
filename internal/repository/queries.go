package repository

// users
const (
	registerUser  = `INSERT INTO users(username,password_hash,is_admin)VALUES ($1,$2,$3)`
	getUserByName = `SELECT id,username,password_hash,is_admin FROM users WHERE username = $1`
	addToCart     = `INSERT INTO user_carts(user_id,product_id,quantity)VALUES ($1,$2,$3)`
	getUserCart   = `SELECT u.id as id,product_name,quantity as qty,selling_price,promo_price,
       u.created_at as created_at,u.updated_at as updated_at FROM user_carts u JOIN products p ON u.product_id = p.id WHERE u.user_id = $1`
)

// product
const (
	getListCategories    = `SELECT id,category_name from product_categories`
	getProductByCategory = `SELECT p.id as id,product_name,category_name as category,qty,selling_price,promo_price FROM products p JOIN product_categories pc
							ON p.category_id = pc.id WHERE p.category_id = $1`
	getProducts = `SELECT p.id as id,product_name,category_name as category,qty,selling_price,promo_price FROM products p JOIN product_categories pc
							ON p.category_id = pc.id`
	getProductDetail = `SELECT p.id as id,product_name,category_name as category,description,qty,selling_price,promo_price,updated_at,created_at FROM products p JOIN product_categories pc
							ON p.category_id = pc.id WHERE p.id = $1`
	UpdateProductQuantity = `UPDATE products SET qty= $1,updated_at=$2 WHERE id =$3`
	insertNewProduct      = `INSERT INTO products(product_name,category_id,description,qty,selling_price,promo_price) VALUES ($1,$2,$3,$4,$5,$6)`
	updateProduct         = `UPDATE products SET product_name=$1,category_id=$2,description=$3,qty=$4,selling_price=$5,promo_price=$6,updated_at=$7 WHERE id = $8`
)
