package database

import (
	"github.com/otaviobelfort/go/product_service/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error("ERROR ...", err)
	}
	err = db.AutoMigrate(&entity.Product{})
	if err != nil {
		return
	}
	product, err := entity.NewProduct("Product 1", 10)
	if err != nil {
		t.Error(err)
	}
	productDB := NewProduct(db)
	assert.NoError(t, productDB.Create(product))
	assert.NotEmpty(t, product.ID)

}
