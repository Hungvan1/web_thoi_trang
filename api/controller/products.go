package controller

import (
	"D/web-thoitrang/api/constant"
	"D/web-thoitrang/api/model"
	"D/web-thoitrang/api/repository"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	productRepo *repository.Product
}

type IDRequest struct {
	ID int `uri:"id" binding:"required,numeric,gt=0"`
}

func NewProduct(productRepo *repository.Product) *Product {
	return &Product{
		productRepo: productRepo,
	}
}

func (p *Product) InitRoutes(r *gin.Engine) {
	r.GET("/products", p.GetProducts)
	r.GET("/products/:id", p.GetProduct)
	r.POST("/products", p.CreateProduct)
	r.PUT("/products/:id", p.UpdateProduct)
	r.DELETE("/products/:id", p.DeleteProduct)
}

func (p *Product) GetProducts(c *gin.Context) {
	products, err := p.productRepo.GetAll(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, constant.Response{
		Data: products,
	})
}

func (p *Product) GetProduct(c *gin.Context) {
	request := &IDRequest{}
	err := c.ShouldBindUri(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, constant.Response{
			Error: &constant.ErrorResponse{
				Message: constant.ErrInvalidRequest,
			},
		})
		return
	}

	product, err := p.productRepo.GetByID(c, request.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, constant.Response{
				Error: &constant.ErrorResponse{
					Message: constant.ErrProductNotFound,
				},
			})
			return
		}
		log.Println("Get product error: ", err)
		c.JSON(http.StatusInternalServerError, constant.Response{
			Error: &constant.ErrorResponse{
				Message: constant.ErrInternalServerError,
			},
		})
		return
	}

	c.JSON(http.StatusOK, constant.Response{
		Data: product,
	})
}

func (p *Product) CreateProduct(c *gin.Context) {
	product := &model.Product{}
	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(http.StatusBadRequest, constant.Response{
			Error: &constant.ErrorResponse{
				Message: constant.ErrInvalidRequest,
			},
		})
		return
	}

	if err := p.productRepo.Create(c, product); err != nil {
		log.Println("Create product error: ", err)
		c.JSON(http.StatusInternalServerError, constant.Response{
			Error: &constant.ErrorResponse{
				Message: constant.ErrInternalServerError,
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": p})
}

// PUT update product
func (p *Product) UpdateProduct(c *gin.Context) {
	request := &IDRequest{}
	err := c.ShouldBindUri(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, constant.Response{
			Error: &constant.ErrorResponse{
				Message: constant.ErrInvalidRequest,
			},
		})
	}

	product := &model.Product{}
	if err = c.ShouldBindJSON(product); err != nil {
		c.JSON(http.StatusBadRequest, constant.Response{
			Error: &constant.ErrorResponse{
				Message: constant.ErrInvalidRequest,
			},
		})
		return
	}

	_, err = p.productRepo.GetByID(c, request.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, constant.Response{
				Error: &constant.ErrorResponse{
					Message: constant.ErrProductNotFound,
				},
			})
			return
		}

		log.Println("Update product error: ", err)
		c.JSON(http.StatusInternalServerError, constant.Response{
			Error: &constant.ErrorResponse{
				Message: constant.ErrInternalServerError,
			},
		})
		return
	}

	product.ID = request.ID
	err = p.productRepo.Update(c, product)
	if err != nil {
		log.Println("Update product error: ", err)
		c.JSON(http.StatusInternalServerError, constant.Response{
			Error: &constant.ErrorResponse{
				Message: constant.ErrInternalServerError,
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DELETE product
func (p *Product) DeleteProduct(c *gin.Context) {
	request := &IDRequest{}
	err := c.ShouldBindUri(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, constant.Response{
			Error: &constant.ErrorResponse{
				Message: constant.ErrInvalidRequest,
			},
		})
	}

	err = p.productRepo.Delete(c, request.ID)
	if err != nil {
		log.Println("Delete product error: ", err)
		c.JSON(http.StatusInternalServerError, constant.Response{
			Error: &constant.ErrorResponse{
				Message: constant.ErrInternalServerError,
			},
		})
		return
	}

	c.JSON(http.StatusOK, constant.Response{
		Message: "Product deleted successfully",
	})
}
