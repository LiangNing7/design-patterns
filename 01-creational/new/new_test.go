package new

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductGetName(t *testing.T) {
	assert := assert.New(t)

	product := NewProduct("Laptop", 20.01)
	assert.Equal("Laptop", product.GetName())
}

func TestProductGetPrice(t *testing.T) {
	assert := assert.New(t)

	product := NewProduct("Laptop", 20.01)
	assert.Equal(20.01, product.GetPrice())
}
