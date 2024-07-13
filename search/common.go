package search

import "errors"

var ErrIsNotSorted = errors.New("input is not sorted")
var ErrNotFound = errors.New("target doesn't exist in the given array")

type testExpectation struct {
	err    error
	result int
}

type searchTest struct {
	input    []int
	target   int
	expected testExpectation
}
