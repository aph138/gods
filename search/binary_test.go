package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var binaryTests = []searchTest{
	{
		input:  []int{},
		target: 2,
		expected: testExpectation{
			err:    ErrNotFound,
			result: -1,
		},
	},
	{
		input:  []int{1},
		target: 1,
		expected: testExpectation{
			err:    nil,
			result: 0,
		},
	},
	{
		input:  []int{1},
		target: 2,
		expected: testExpectation{
			err:    ErrNotFound,
			result: -1,
		},
	},
	{
		input:  []int{1, 2, 3},
		target: 2,
		expected: testExpectation{
			err:    nil,
			result: 1,
		},
	},
	{
		input:  []int{4, 2, 3, 1, 9, 7},
		target: 3,
		expected: testExpectation{
			err:    ErrIsNotSorted,
			result: -1,
		},
	},
	{
		input:  []int{4, 7, 9, 10, 15, 19, 21},
		target: 21,
		expected: testExpectation{
			err:    nil,
			result: 6,
		},
	},
	{
		input:  []int{4, 7, 9, 10, 15, 19, 21},
		target: 4,
		expected: testExpectation{
			err:    nil,
			result: 0,
		},
	},
}

func TestBinarySearch(t *testing.T) {
	assert := assert.New(t)
	for _, c := range binaryTests {
		result, err := BinarySearch(c.input, c.target)
		assert.Equal(c.expected.err, err)
		assert.Equal(c.expected.result, result)
	}
}
