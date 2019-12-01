package counter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent_of_code_2019/day_1/counter"
)

func TestCountFuel(t *testing.T) {
	testCases := []struct{
		input []int
		expected int
	} {
		{
			input: []int{12, 14, 1969, 100756},
			expected: 34241,
		},
	}

	c := counter.NewCounter()

	for _, tc := range testCases {
		assert.Equal(t, tc.expected,  c.CountFuel(tc.input))
	}
}

func TestCountExtraFuel(t *testing.T) {
	testCases := []struct{
		input []int
		expected int
	} {
		{
			input: []int{12, 14, 1969, 100756},
			expected: 51316,
		},
		{
			input: []int{12, 14},
			expected: 4,
		},
		{
			input: []int{12},
			expected: 2,
		},
		{
			input: []int{1969, 100756},
			expected: 51312,
		},
		{
			input: []int{1969},
			expected: 966,
		},
		{
			input: []int{100756},
			expected: 50346,
		},
	}

	c := counter.NewCounter()

	for _, tc := range testCases {
		assert.Equal(t, tc.expected,  c.CountExtraFuel(tc.input))
	}
}
