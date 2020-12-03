package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	numbers := parseFile()

	log.Println("Part one:", partOne(numbers))

	log.Println("Part two:", partTwo(numbers))
}

func parseFile() []int {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers
}

// partOne - find the two entries that sum to 2020 and then multiply those two numbers together.
func partOne(numbers []int) int {
	target := 2020

	c := make(chan [2]int)
	findTwoAddends(numbers, target, c)

	result := <-c
	return result[0] * result[1]
}

// partTwo - find three numbers in your expense report that meet the same criteria.
func partTwo(numbers []int) int {
	target := 2020

	c := make(chan [3]int)
	findThreeAddends(numbers, target, c)

	result := <-c
	return result[0] * result[1] * result[2]
}

func findTwoAddends(numbers []int, targetSum int, addends chan<- [2]int) {
	for i, v := range numbers[:len(numbers)-1] {
		go func(addend int, numbers []int, targetSum int) {
			for _, v := range numbers {
				sum := addend + v
				if sum == targetSum {
					addends <- [2]int{addend, v}
				}
			}
		}(v, numbers[i+1:], targetSum)
	}
}

func findThreeAddends(numbers []int, targetSum int, addends chan<- [3]int) {
	for i, v := range numbers[:len(numbers)-1] {
		go func(addend int, numbers []int, targetSum int) {
			for i, v := range numbers {
				go func(addend1 int, addend2 int, numbers []int, targetSum int) {
					for _, v := range numbers {
						sum := addend1 + addend2 + v
						if sum == targetSum {
							addends <- [3]int{addend1, addend2, v}
						}
					}
				}(addend, v, numbers[i+1:], targetSum)
			}
		}(v, numbers[i+1:], targetSum)
	}
}
