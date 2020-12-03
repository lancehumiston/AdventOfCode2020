package main

import (
	"testing"
)

type test struct {
	name     string
	input    []Record
	expected int
}

var partOneTests = []test{
	{
		name: "Part 1 - AoC Example #1",
		input: []Record{
			Record{
				Min:      1,
				Max:      3,
				Letter:   rune('a'),
				Password: "abcde",
			},
			Record{
				Min:      1,
				Max:      3,
				Letter:   rune('b'),
				Password: "cdefg",
			},
			Record{
				Min:      2,
				Max:      9,
				Letter:   rune('c'),
				Password: "ccccccccc",
			},
		},
		expected: 2,
	},
}

var partTwoTests = []test{
	{
		name: "Part 2 - AoC Example #1",
		input: []Record{
			Record{
				Min:      1,
				Max:      3,
				Letter:   rune('a'),
				Password: "abcde",
			},
			Record{
				Min:      1,
				Max:      3,
				Letter:   rune('b'),
				Password: "cdefg",
			},
			Record{
				Min:      2,
				Max:      9,
				Letter:   rune('c'),
				Password: "ccccccccc",
			},
		},
		expected: 1,
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
