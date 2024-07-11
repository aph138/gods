package sorts

type sortTest struct {
	input    []int
	expected []int
}

var tests = []sortTest{
	{
		input:    []int{},
		expected: []int{},
	},
	{
		input:    []int{1},
		expected: []int{1},
	},
	{
		input:    []int{1, 1},
		expected: []int{1, 1},
	},
	{
		input:    []int{1, 2, 3},
		expected: []int{1, 2, 3},
	},
	{
		input:    []int{3, 2, 1},
		expected: []int{1, 2, 3},
	},
	{
		input:    []int{5, 2, 1},
		expected: []int{1, 2, 5},
	},
}
