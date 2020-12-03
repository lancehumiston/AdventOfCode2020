package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	grid := parseFile()

	log.Println("Part one:", partOne(grid))

	log.Println("Part two:", partTwo(grid))
}

// partOne - count all the trees('#') you would encounter for the slope right 3, down 1
func partOne(grid [][]byte) (count int) {
	const tree = '#'

	height := len(grid)
	if height == 0 {
		log.Panic("empty grid is invalid")
	}
	width := len(grid[0])

	for x, y := 0, 0; y < height; x, y = x+3, y+1 {
		if x >= width {
			x = x - width
		}

		if grid[y][x] == tree {
			count++
		}
	}

	return
}

// partTwo - count all the trees('#') you would encounter for each slope, then multiply them together
func partTwo(grid [][]byte) (count int) {
	const tree = '#'

	// collection of slopes [{run}, {-rise}]
	slopes := [][2]int{
		[2]int{1, 1},
		[2]int{3, 1},
		[2]int{5, 1},
		[2]int{7, 1},
		[2]int{1, 2},
	}

	height := len(grid)
	if height == 0 {
		log.Panic("empty grid is invalid")
	}
	width := len(grid[0])

	wg := sync.WaitGroup{}
	wg.Add(len(slopes))

	for _, s := range slopes {
		go func(slope [2]int) {
			defer wg.Done()

			c := 0
			for x, y := 0, 0; y < height; x, y = x+slope[0], y+slope[1] {
				if x >= width {
					x = x - width
				}

				if grid[y][x] == tree {
					c++
				}
			}

			if count == 0 {
				count = c
			} else {
				count = count * c
			}

		}(s)
	}

	wg.Wait()

	return
}

func parseFile() [][]byte {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()

	grid := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	fmt.Println("elapsed", time.Since(start))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}
