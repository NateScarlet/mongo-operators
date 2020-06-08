package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeOperators(t *testing.T) {
	res := MergeOperators(M{"a": 1}, M{"b": 2}, M{"a": 3})
	assert.Equal(t, M{"a": 3, "b": 2}, res)
}
