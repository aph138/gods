package sorts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubble(t *testing.T) {
	assert := assert.New(t)
	for _, c := range tests {
		BubbleSort(c.input)
		assert.Equal(c.expected, c.input)
	}
}
