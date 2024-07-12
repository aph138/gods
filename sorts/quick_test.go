package sorts

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	assert := assert.New(t)

	for _, c := range tests {
		result := slices.Clone(c.input)
		QuickSort(result)
		assert.Equal(c.expected, result)
	}

}
