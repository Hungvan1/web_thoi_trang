package repository

import (
	"D/web-thoitrang/api/model"
	"context"

	"gorm.io/gorm"
)

type Product struct {
	db *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		db: db,
	}
}

func (p *Product) GetAll(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	result := p.db.WithContext(ctx).Find(&products)
	return products, result.Error
}

func (p *Product) GetByID(ctx context.Context, id int) (*model.Product, error) {
	product := &model.Product{}
	result := p.db.WithContext(ctx).First(product, id)
	return product, result.Error
}

func (p *Product) GetByIDs(ctx context.Context, ids []int) ([]*model.Product, error) {
	var products []*model.Product
	result := p.db.WithContext(ctx).Find(&products, ids)
	return products, result.Error
}

func (p *Product) Create(ctx context.Context, product *model.Product) error {
	return p.db.WithContext(ctx).Create(product).Error
}

func (p *Product) Update(ctx context.Context, product *model.Product) error {
	return p.db.WithContext(ctx).Updates(product).Error
}

func (p *Product) Delete(ctx context.Context, id int) error {
	return p.db.WithContext(ctx).Delete(&model.Product{}, id).Error
}
