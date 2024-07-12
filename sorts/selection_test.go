package sorts

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelection(t *testing.T) {
	assert := assert.New(t)
	for _, c := range tests {
		result := slices.Clone(c.input)
		SelectionSort(result)
		assert.Equal(c.expected, result)
	}
}
