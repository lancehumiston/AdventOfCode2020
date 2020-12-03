package main

import (
	"testing"
)

type test struct {
	name     string
	input    [][]byte
	expected int
}

var input = [][]byte{
	[]byte("..##......."),
	[]byte("#...#...#.."),
	[]byte(".#....#..#."),
	[]byte("..#.#...#.#"),
	[]byte(".#...##..#."),
	[]byte("..#.##....."),
	[]byte(".#.#.#....#"),
	[]byte(".#........#"),
	[]byte("#.##...#..."),
	[]byte("#...##....#"),
	[]byte(".#..#...#.#"),
}

var partOneTests = []test{
	{
		name:     "Part 1 - AoC Example #1",
		input:    input,
		expected: 7,
	},
}

var partTwoTests = []test{
	{
		name:     "Part 2 - AoC Example #1",
		input:    input,
		expected: 336,
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
