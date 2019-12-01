package reader_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent_of_code_2019/day_1/reader"
)

func TestRead(t *testing.T) {
	testCases := []struct{
		fileName string
		expected []int
	} {
		{
			fileName: "testdata/test.txt",
			expected: []int{12, 14, 1969, 100756},
		},
	}

	r := reader.NewFileReader()

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, r.Read(tc.fileName))
	}
}
