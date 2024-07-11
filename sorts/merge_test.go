package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	assert := assert.New(t)
	for _, c := range tests {
		assert.Equal(c.expected, MergeSort(c.input))
	}
}
