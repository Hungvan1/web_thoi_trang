package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Number      int     `json:"number"`
	Detail      string  `json:"detail"`
	Status      string  `json:"status"`
	Size        string  `json:"size"`
	Gender      string  `json:"gender"`
	Color       string  `json:"color"`
	CateID      int     `json:"cate_id"`
	UserID      int     `json:"user_id"`
	Image       string  `json:"image"`
}

type User struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Category struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type Order struct {
	OrderID     int     `json:"order_id"`
	OrderDate   string  `json:"order_date"`
	ShipAddress string  `json:"ship_address"`
	UserID      int     `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
}
type OrderItem struct {
	OrderItemID int     `json:"order_item_id"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	OrderID     int     `json:"order_id"`
	ProductID   int     `json:"product_id"`
}

type Review struct {
	ReviewID  int    `json:"review_id"`
	Text      string `json:"text"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	OrderID   int    `json:"order_id"`
}

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/clothing_shop")
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}
	log.Println("Connected to MySQL successful!")
}

func main() {
	ConnectDB()

	r := gin.Default()

	r.GET("/products", GetProducts)
	r.GET("/products/:id", GetProduct)
	r.POST("/products", CreateProduct)
	r.PUT("/products/:id", UpdateProduct)
	r.DELETE("/products/:id", DeleteProduct)

	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUser)
	r.POST("/users", CreateUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	r.GET("/categories", GetCategories)
	r.GET("/categories/:id", GetCategory)
	r.POST("/categories", CreateCategory)
	r.PUT("/categories/:id", UpdateCategory)
	r.DELETE("/categories/:id", DeleteCategory)

	r.GET("/orders", GetOrders)
	r.GET("/orders/:id", GetOrder)
	r.POST("/orders", CreateOrder)
	r.PUT("/orders/:id", UpdateOrder)
	r.DELETE("/orders/:id", DeleteOrder)

	r.GET("/order-items", GetOrderItems)
	r.GET("/order-items/:id", GetOrderItem)
	r.POST("/order-items", CreateOrderItem)
	r.PUT("/order-items/:id", UpdateOrderItem)
	r.DELETE("/order-items/:id", DeleteOrderItem)

	r.GET("/reviews", GetReviews)
	r.GET("/reviews/:id", GetReview)
	r.POST("/reviews", CreateReview)
	r.PUT("/reviews/:id", UpdateReview)
	r.DELETE("/reviews/:id", DeleteReview)

	r.Run(":8080")
}

// GET all products
func GetProducts(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ProductID, &p.ProductName, &p.Price, &p.Number,
			&p.Detail, &p.Status, &p.Size, &p.Gender, &p.Color, &p.CateID, &p.UserID, &p.Image); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, p)
	}
	c.JSON(http.StatusOK, products)
}

// GET one product
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var p Product
	err := DB.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(
		&p.ProductID, &p.ProductName, &p.Price, &p.Number, &p.Detail,
		&p.Status, &p.Size, &p.Gender, &p.Color, &p.CateID, &p.UserID, &p.Image,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, p)
}

// POST create product
func CreateProduct(c *gin.Context) {
	var p Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.Exec(`INSERT INTO products 
		(product_name, price, number, detail, status, size, gender, color, category_id, user_id, image)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.ProductName, p.Price, p.Number, p.Detail, p.Status, p.Size, p.Gender, p.Color, p.CateID, p.UserID, p.Image)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}

// PUT update product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var p Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.Exec(`UPDATE products 
		SET product_name=?, price=?, number=?, detail=?, status=?, size=?, gender=?, color=?, category_id=?, user_id=?, image=? 
		WHERE id=?`,
		p.ProductName, p.Price, p.Number, p.Detail, p.Status, p.Size, p.Gender, p.Color, p.CateID, p.UserID, p.Image, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DELETE product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	_, err := DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// GET all users
func GetUsers(c *gin.Context) {
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.UserID, &u.Name, &u.Phone, &u.Address, &u.Username, &u.Password, &u.Role); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, u)
	}
	c.JSON(http.StatusOK, users)
}

// GET one user
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var u User
	err := DB.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(
		&u.UserID, &u.Name, &u.Phone, &u.Address, &u.Username, &u.Password, &u.Role,
	)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}

// POST create user
func CreateUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.Exec(`INSERT INTO users (name, phone, address, username, password, role)
		VALUES (?, ?, ?, ?, ?, ?)`,
		u.Name, u.Phone, u.Address, u.Username, u.Password, u.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// PUT update user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := DB.Exec(`UPDATE users SET name=?, phone=?, address=?, username=?, password=?, role=? WHERE id=?`,
		u.Name, u.Phone, u.Address, u.Username, u.Password, u.Role, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DELETE user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
