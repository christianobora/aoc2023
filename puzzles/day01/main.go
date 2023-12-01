package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/christianobora/aoc2023/pkg/files"
	"github.com/christianobora/aoc2023/pkg/ints"
)

func main() {
	fmt.Println("Part 1: ", getCalibrationValues("input.txt", false))
	fmt.Println("Part 2: ", getCalibrationValues("input.txt", true))
}

func getCalibrationValues(input string, includeSpelled bool) int {
	lines := files.ReadLines(input)
	totalCalibrationValue := 0
	for _, line := range lines {
		firstDigit := getFirstDigit(line, includeSpelled)
		lastDigit := getLastDigit(line, includeSpelled)
		combined := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
		totalCalibrationValue += ints.FromString(combined)
	}
	return totalCalibrationValue
}

func getFirstDigit(s string, includeSpelled bool) int {
	for i := 0; i < len(s); i++ {
		if i, ok := isDigit(s[i]); ok {
			return i
		}
		if includeSpelled {
			if i, ok := spelledDigit(s[i:]); ok {
				return i
			}
		}
	}
	panic("No digit found in " + s)
}

func getLastDigit(s string, includeSpelled bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if i, ok := isDigit(s[i]); ok {
			return i
		}
		if includeSpelled {
			if i, ok := reverseSpelledDigit(s[:i+1]); ok {
				return i
			}
		}
	}
	panic("No digit found in " + s)
}

func isDigit(b byte) (int, bool) {
	if unicode.IsDigit(rune(b)) {
		return int(b - '0'), true
	}
	return 0, false
}

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func spelledDigit(s string) (int, bool) {
	for k, v := range digitMap {
		if strings.HasPrefix(s, k) {
			return v, true
		}
	}
	return -1, false
}

func reverseSpelledDigit(s string) (int, bool) {
	for k, v := range digitMap {
		if strings.HasSuffix(s, k) {
			return v, true
		}
	}
	return -1, false
}
