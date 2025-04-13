package objectpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObjectPool(t *testing.T) {
	assert := assert.New(t)

	pool := NewObjectPool(3)

	object1 := pool.AcquireObject()
	assert.Equal(0, object1.ID)

	object2 := pool.AcquireObject()
	assert.Equal(1, object2.ID)

	pool.ReleaseObject(object1)
	assert.Equal(0, object1.ID)

	object3 := pool.AcquireObject()
	assert.Equal(2, object3.ID)
}
