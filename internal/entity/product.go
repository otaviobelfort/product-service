package entity

import (
	"errors"
	"github.com/otaviobelfort/go/product_service/pkg/entity"
	"time"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil

}

func (product *Product) Validate() error {
	if product.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(product.ID.String()); err != nil {
		return ErrInvalidID
	}
	if product.Name == "" {
		return ErrNameIsRequired
	}
	if product.Price <= 0 {
		return ErrInvalidPrice
	}
	return nil

}
