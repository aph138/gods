package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	assert := assert.New(t)
	s := NewStack(1, 5)
	// 1,5
	assert.Equal(2, s.Len())
	result, err := s.Peek()
	assert.NoError(err)
	assert.Equal(5, result)
	// 1
	result, err = s.Pop()
	assert.NoError(err)
	assert.Equal(5, result)
	assert.Equal(1, s.Len())
	result, err = s.Peek()
	assert.NoError(err)
	assert.Equal(1, result)
	// empty
	result, err = s.Pop()
	assert.NoError(err)
	assert.Equal(1, result)
	_, err = s.Pop()
	assert.ErrorIs(err, ErrEmpty)
	_, err = s.Peek()
	assert.ErrorIs(err, ErrEmpty)
	// 6
	s.Add(6)
	assert.Equal(1, s.Len())
	result, err = s.Peek()
	assert.NoError(err)
	assert.Equal(6, result)

}
