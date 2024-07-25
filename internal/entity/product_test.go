package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Celular", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Celular", product.Name)
	assert.Equal(t, 10, product.Price)
}
func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 10)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}
func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Celular", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}
func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Celular", -10)
	//assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}
