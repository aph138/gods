package sorts

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertion(t *testing.T) {
	assert := assert.New(t)
	for _, c := range tests {
		result := slices.Clone(c.input)
		InsertionSort(result)
		assert.Equal(c.expected, result)
	}
}
