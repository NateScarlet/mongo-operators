package aggregation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnless(t *testing.T) {
	var res M
	res = Unless(M{"$eq": A{1, 1}}, true)
	assert.Equal(t, M{"$eq": A{1, 1}}, res)
	res = Unless(M{"$eq": A{1, 1}}, false)
	assert.Equal(t, M{"$not": M{"$eq": A{1, 1}}}, res)
}
