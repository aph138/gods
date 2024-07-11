package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAndDelete(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		input    *List[int]
		expected *List[int]
	}{
		{
			input:    NewList[int]().Add(2),
			expected: NewList(2),
		},
		{
			input:    NewList[int](1, 2, 3, 4, 5).RemoveAt(1),
			expected: NewList(1, 3, 4, 5),
		},
	}

	for _, c := range cases {
		assert.Equal(c.expected, c.input)
	}

}
func TestLength(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		input    *List[int]
		expected int
	}{
		{
			input:    NewList(1, 2, 3).RemoveAt(0).RemoveAt(0),
			expected: 1,
		},
	}
	for _, c := range cases {
		assert.Equal(c.expected, c.input.Len())
	}
}
