package aoc20

import (
	"regexp"
	"slices"

	. "github.com/roidaradal/aoc-go/aoc"
	"github.com/roidaradal/fn"
)

func Day04() Solution {
	passports := data04(true)
	count1, count2 := 0, 0
	for _, p := range passports {
		// Part 1
		if hasAllKeys(p) {
			count1 += 1
		}

		// Part 2
		if isValidPassport(p) {
			count2 += 1
		}
	}
	return NewSolution(count1, count2)
}

func data04(full bool) []Passport {
	passports := make([]Passport, 0)
	p := make(Passport)
	for _, line := range ReadRawLines(20, 4, full, true) {
		if line == "" {
			passports = append(passports, p)
			p = make(Passport)
		} else {
			for _, pair := range fn.SpaceSplit(line) {
				parts := fn.CleanSplit(pair, ":")
				k, v := parts[0], parts[1]
				p[k] = v
			}
		}
	}
	passports = append(passports, p)
	return passports
}

type Passport = map[string]string

var passportKeys = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func hasAllKeys(p Passport) bool {
	return fn.All(passportKeys, func(key string) bool {
		return fn.HasKey(p, key)
	})
}

func isValidPassport(p Passport) bool {
	hclPattern := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	pidPattern := regexp.MustCompile(`^[0-9]{9}$`)
	eclOptions := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	ok := 0
	for k, v := range p {
		switch k {
		case "byr":
			byr := fn.ParseInt(v)
			if 1920 <= byr && byr <= 2002 {
				ok += 1
			}
		case "iyr":
			iyr := fn.ParseInt(v)
			if 2010 <= iyr && iyr <= 2020 {
				ok += 1
			}
		case "eyr":
			eyr := fn.ParseInt(v)
			if 2020 <= eyr && eyr <= 2030 {
				ok += 1
			}
		case "hgt":
			n := len(v)
			hgt, unit := fn.ParseInt(v[:n-2]), v[n-2:]
			if unit == "cm" && 150 <= hgt && hgt <= 193 {
				ok += 1
			}
			if unit == "in" && 59 <= hgt && hgt <= 76 {
				ok += 1
			}
		case "hcl":
			if hclPattern.MatchString(v) {
				ok += 1
			}
		case "ecl":
			if slices.Contains(eclOptions, v) {
				ok += 1
			}
		case "pid":
			if pidPattern.MatchString(v) {
				ok += 1
			}
		}
	}
	return ok == len(passportKeys)
}
