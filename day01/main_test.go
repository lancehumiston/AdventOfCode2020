package main

import (
	"testing"
)

type test struct {
	name     string
	input    []int
	expected int
}

var partOneTests = []test{
	{
		name:     "Part 1 - AoC Example #1",
		input:    []int{1721, 979, 366, 299, 675, 1456},
		expected: 514579,
	},
}

var partTwoTests = []test{
	{
		name:     "Part II - AoC Example #1",
		input:    []int{1721, 979, 366, 299, 675, 1456},
		expected: 241861950,
	},
}

func TestPartOne(t *testing.T) {
	for _, tc := range partOneTests {
		actual := partOne(tc.input)

		if actual != tc.expected {
			t.Fatalf("%s expected: %v, actual: %v", tc.name, tc.expected, actual)
		}
	}
}

func TestPartTwo(t *testing.T) {
	for _, tc := range partTwoTests {
		actual := partTwo(tc.input)

		if actual != tc.expected {
			t.Fatalf("%s expected: %v, actual: %v", tc.name, tc.expected, actual)
		}
	}
}
