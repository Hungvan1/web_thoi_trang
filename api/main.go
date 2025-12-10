package main

import (
	"D/web-thoitrang/api/controller"
	"D/web-thoitrang/api/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	UserID   int    `gorm:"primary_key;column:id" json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `son:"role"`
}

type Category struct {
	CategoryID   int    `gorm:"primary_key;column:id" json:"id"`
	CategoryName string `json:"category_name"`
}

type Order struct {
	OrderID     int     `gorm:"primary_key;column:id" json:"id"`
	OrderDate   string  `json:"order_date"`
	ShipAddress string  `json:"ship_address"`
	UserID      int     `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
}

type OrderItem struct {
	OrderItemID int     `gorm:"primary_key;column:id" json:"id"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	OrderID     int     `json:"order_id"`
	ProductID   int     `json:"product_id"`
}

type Review struct {
	ReviewID  int    `gorm:"primary_key;column:id" json:"review_id"`
	Text      string `json:"text"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	OrderID   int    `json:"order_id"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func ConnectDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/clothing_shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	log.Println("Connected to database")
	return db
}

func main() {
	db := ConnectDB()
	r := gin.Default()

	productRepo := repository.NewProduct(db)
	productCtrl := controller.NewProduct(productRepo)
	productCtrl.InitRoutes(r)

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

	r.GET("/order_items", GetOrderItems)
	r.GET("/order_items/:id", GetOrderItem)
	r.POST("/order_items", CreateOrderItem)
	r.PUT("/order_items/:id", UpdateOrderItem)
	r.DELETE("/order_items/:id", DeleteOrderItem)

	r.GET("/reviews", GetReviews)
	r.GET("/reviews/:id", GetReview)
	r.POST("/reviews", CreateReview)
	r.PUT("/reviews/:id", UpdateReview)
	r.DELETE("/reviews/:id", DeleteReview)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}

// GET all users
func GetUsers(c *gin.Context) {
	var users []User
	if result := DB.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GET one user
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var u User
	if result := DB.First(&u, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
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
	if result := DB.Create(&u); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": u})
}

// PUT update user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser User
	if result := DB.First(&existingUser, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	if result := DB.Model(&existingUser).Updates(u); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DELETE user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if result := DB.Delete(&User{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GET all categories
func GetCategories(c *gin.Context) {
	var categories []Category
	if result := DB.Find(&categories); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GET one category
func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var cat Category

	if result := DB.First(&cat, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, cat)
}

// POST create category
func CreateCategory(c *gin.Context) {
	var cat Category
	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Create(&cat); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully", "category": cat})
}

// PUT update category
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var cat Category

	if err := c.ShouldBindJSON(&cat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Model(&Category{}).Where("id = ?", id).Update("CategoryName", cat.CategoryName); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result := DB.Model(&Category{}).Where("id = ?", id).Update("CategoryName", cat.CategoryName); result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

// DELETE category
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	if result := DB.Delete(&Category{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// GET all orders
func GetOrders(c *gin.Context) {
	var orders []Order
	if result := DB.Find(&orders); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GET one order
func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var o Order

	if result := DB.First(&o, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, o)
}

// POST create order
func CreateOrder(c *gin.Context) {
	var o Order

	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Create(&o); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": o})
}

// PUT update order
func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var o Order

	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Model(&Order{}).Where("id = ?", id).Updates(o); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result := DB.Model(&Order{}).Where("id = ?", id).Updates(o); result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

// DELETE order
func DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	if result := DB.Delete(&Order{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// GET all orderitem
func GetOrderItems(c *gin.Context) {
	var items []OrderItem
	if result := DB.Find(&items); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// Get one orderitem
func GetOrderItem(c *gin.Context) {
	id := c.Param("id")
	var it OrderItem

	if result := DB.First(&it, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Order item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, it)
}

// POST orderitem
func CreateOrderItem(c *gin.Context) {
	var it OrderItem

	if err := c.ShouldBindJSON(&it); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Create(&it); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order item created successfully", "order_item": it})
}

// PUT orderitem
func UpdateOrderItem(c *gin.Context) {
	id := c.Param("id")
	var it OrderItem

	if err := c.ShouldBindJSON(&it); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Model(&OrderItem{}).Where("id = ?", id).Updates(it); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result := DB.Model(&OrderItem{}).Where("id = ?", id).Updates(it); result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Order item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order item updated successfully"})
}

// DELETE orderitem
func DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")

	if result := DB.Delete(&OrderItem{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order item deleted successfully"})
}

// GET all reviews
func GetReviews(c *gin.Context) {
	var reviews []Review
	if result := DB.Find(&reviews); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

// GET one review
func GetReview(c *gin.Context) {
	id := c.Param("id")
	var r Review

	if result := DB.First(&r, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Review not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, r)
}

// POST review
func CreateReview(c *gin.Context) {
	var r Review

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Create(&r); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Review created successfully", "review": r})
}

// PUT review
func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var r Review

	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := DB.Model(&Review{}).Where("review_id = ?", id).Updates(r); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result := DB.Model(&Review{}).Where("review_id = ?", id).Updates(r); result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Review not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review updated successfully"})
}

// DELETE review
func DeleteReview(c *gin.Context) {
	id := c.Param("id")

	if result := DB.Delete(&Review{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}
