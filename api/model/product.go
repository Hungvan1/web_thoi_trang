package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type ProductStatus int

const (
	ProductStatusAvailable ProductStatus = iota
	ProductStatusUnavailable
)

var productValues = []string{"AVAILABLE", "UNAVAILABLE"}

type Product struct {
	ID          int            `gorm:"primary_key;column:id" json:"id"`
	ProductName string         `json:"product_name"`
	Price       float64        `json:"price"`
	Number      int            `json:"number"`
	Detail      string         `json:"detail"`
	Status      *ProductStatus `json:"status"`
	Size        string         `json:"size"`
	Gender      string         `json:"gender" gorm:"column:gender"`
	Color       string         `json:"color" gorm:"column:color"`
	CategoryID  int            `json:"category_id" gorm:"column:category_id"`
	UserID      int            `json:"user_id"`
	Image       string         `json:"image"`
}

func (p *Product) TableName() string {
	return "products"
}

func (s *ProductStatus) String() string {
	return productValues[*s]
}

func ProductStatusFromString(s string) (ProductStatus, error) {
	for i, v := range productValues {
		if v == s {
			return ProductStatus(i), nil
		}
	}
	return ProductStatus(0), errors.New("invalid ProductStatus string")
}

func (s *ProductStatus) Value() (driver.Value, error) {
	return s.String(), nil
}

func (s *ProductStatus) Scan(value interface{}) (err error) {
	val, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	*s, err = ProductStatusFromString(string(val))
	return
}

func (s *ProductStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", s.String())), nil
}

func (s *ProductStatus) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	str = str[1 : len(str)-1] // Remove quotes
	*s, err = ProductStatusFromString(str)
	return
}
