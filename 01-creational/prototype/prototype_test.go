package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleClone(t *testing.T) {
	assert := assert.New(t)

	circle := &Circle{Type: "Circle"}
	cloneCircle := circle.Clone()

	actual := cloneCircle.GetType()
	assert.Equal(circle.GetType(), actual)
}
