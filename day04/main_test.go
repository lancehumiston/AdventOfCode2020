package main

import (
	"testing"
)

type test struct {
	name     string
	input    []string
	expected int
}

var partOneTests = []test{
	{
		name: "Part 1 - AoC Example #1",
		input: []string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
			"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
			"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
		},
		expected: 2,
	},
}

var partTwoTests = []test{
	{
		name: "Part 2 - AoC Example #1",
		input: []string{
			"eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
			"iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946",
			"hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
			"hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007",
			"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f",
			"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
			"hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022",
			"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
		},
		expected: 4,
	},
	{
		name: "All invalid examples",
		input: []string{
			"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:2003 hcl:#623a2f",
			"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:190in",
			"hcl:#888785 hgt:190 byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022",
			"iyr:2010 hgt:158cm hcl:#123abz ecl:blu byr:1944 eyr:2021 pid:093154719",
			"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:0123456789 hcl:#a97842 hgt:165cm",
			"hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:wat eyr:2022",
			"iyr:2010 hgt:158cm hcl:123abc ecl:blu byr:1944 eyr:2021 pid:093154719",
		},
		expected: 0,
	},
	{
		name: "All valid examples",
		input: []string{
			"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:2002 hcl:#623a2f",
			"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:60in",
			"hcl:#888785 hgt:190cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022",
			"iyr:2010 hgt:158cm hcl:#123abc ecl:blu byr:1944 eyr:2021 pid:093154719",
			"pid:087499704 hgt:74in ecl:brn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f",
			"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:000000001 hcl:#a97842 hgt:165cm",
		},
		expected: 6,
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
