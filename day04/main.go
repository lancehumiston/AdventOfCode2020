package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	passports := parseFile()

	log.Println("Part one:", partOne(passports))

	log.Println("Part two:", partTwo(passports))
}

// partOne - count the number of valid passports - those that have all required fields. Treat cid as optional.
func partOne(passports []string) (count int) {
	for _, v := range passports {
		if hasRequiredFields(v) {
			count++
		}
	}

	return
}

// hasRequiredFields - validates the passport against the required fields
func hasRequiredFields(passport string) bool {
	reqFields := [...]string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for i := range reqFields {
		if !strings.Contains(passport, reqFields[i]) {
			return false
		}
	}

	return true
}

// partTwo - count the passports where all required fields are both present and valid
func partTwo(passports []string) (count int) {
	for _, v := range passports {
		if isValid(v) {
			count++
		}
	}

	return
}

// isValid - validates the passport against the required fields and validation rules
func isValid(passport string) bool {
	fieldValidators := map[string]*regexp.Regexp{
		"byr": regexp.MustCompile(`^19[2-8][0-9]|199[0-9]|200[0-2]$`),                 // four digits; at least 1920 and at most 2002.
		"iyr": regexp.MustCompile(`^201[0-9]|2020$`),                                  // four digits; at least 2010 and at most 2020.
		"eyr": regexp.MustCompile(`^202[0-9]|2030$`),                                  // four digits; at least 2020 and at most 2030.
		"hgt": regexp.MustCompile(`^(1[5-8][0-9]|19[0-3])cm$|^(59|6[0-9]|7[0-6])in$`), // a number followed by either cm or in: If cm, the number must be at least 150 and at most 193., If in, the number must be at least 59 and at most 76.
		"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`),                                   // a # followed by exactly six characters 0-9 or a-f.
		"ecl": regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth){1}$`),               // exactly one of: amb blu brn gry grn hzl oth.
		"pid": regexp.MustCompile(`^\d{9}$`),                                          // a nine-digit number, including leading zeroes.
	}

	for field, validator := range fieldValidators {
		r := regexp.MustCompile(fmt.Sprintf(`%s:([^\s]+)`, field))
		value := r.FindStringSubmatch(passport)
		if len(value) < 2 || !validator.MatchString(value[1]) {
			return false
		}
	}

	return true
}

func parseFile() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	records := make([]string, 1)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			records = append(records, line)
			continue
		}

		records[len(records)-1] += fmt.Sprintf(" %s", line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return records
}
