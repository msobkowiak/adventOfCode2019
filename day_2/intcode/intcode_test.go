package intcode_test

import (
	"testing"

	"advent_of_code_2019/day_2/intcode"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	testCases := []struct{
		input []int
		expected []int
	} {
		{
			input: []int{1,0,0,0,99},
			expected: []int{2,0,0,0,99},
		},
		{
			input: []int{2,3,0,3,99},
			expected: []int{2,3,0,6,99},
		},
		{
			input: []int{2,4,4,5,99,0},
			expected: []int{2,4,4,5,99,9801},
		},
		{
			input: []int{1,1,1,4,99,5,6,0,99},
			expected: []int{30,1,1,4,2,5,6,0,99},
		},
	}

	intCode := intcode.NewIntcode()

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, intCode.Calculate(tc.input))
	}
}

func TestRestore(t *testing.T) {
	testCases := []struct{
		input []int
		expected []int
	} {
		{
			input: []int{1,0,0,0,99},
			expected: []int{1,12,2,0,99},
		},
		{
			input: []int{2,3,0,3,99},
			expected: []int{2,12,2,3,99},
		},
		{
			input: []int{2,4,4,5,99,0},
			expected: []int{2,12,2,5,99,0},
		},
		{
			input: []int{1,1,1,4,99,5,6,0,99},
			expected: []int{1,12,2,4,99,5,6,0,99},
		},
	}

	intCode := intcode.NewIntcode()

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, intCode.Restore(tc.input))
	}
}
