package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Record - password database record
type Record struct {
	Min      int
	Max      int
	Letter   rune
	Password string
}

func main() {
	records := parseFile()

	log.Println("Part one:", partOne(records))

	log.Println("Part two:", partTwo(records))
}

// partOne - How many passwords are valid according to their policies.
// The policy(Min & Max) on the Record indicates the lowest and highest
// number of times a given letter must appear for the password to be valid
func partOne(records []Record) int {
	count := 0

	for _, v := range records {
		if c := strings.Count(v.Password, string(v.Letter)); c >= v.Min && c <= v.Max {
			count++
		}
	}

	return count
}

// partTwo - How many passwords are valid according to their policies.
// The policy(Min & Max) on the Record describes two positions in the
// password, not zero-indexed
func partTwo(records []Record) int {
	count := 0

	for _, v := range records {
		s := byte(v.Letter)
		hasFirstPos := v.Password[v.Min-1] == s
		hasSecondPos := v.Password[v.Max-1] == s

		if hasFirstPos != hasSecondPos {
			count++
		}
	}

	return count
}

func parseFile() []Record {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var records []Record
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var record Record
		line := scanner.Text()
		if _, err := fmt.Sscanf(line, "%d-%d %c: %s", &record.Min, &record.Max, &record.Letter, &record.Password); err != nil {
			log.Println(record)
		}
		records = append(records, record)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return records
}
