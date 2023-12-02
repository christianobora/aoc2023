package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/christianobora/aoc2023/pkg/files"
)

var (
	diceRegex = regexp.MustCompile(`(\d+) (\w+)`)
)

func main() {
	fmt.Println("Part 1: ", part1("input.txt"))
	fmt.Println("Part 2: ", part2("input.txt"))
}

func part1(input string) int {
	part1, _ := solve(input)
	return part1
}

func part2(name string) int {
	_, part2 := solve(name)
	return part2
}

func solve(input string) (int, int) {
	parsedInput := files.Read(input)

	part1, part2 := 0, 0
	for i, s := range strings.Split(strings.TrimSpace(string(parsedInput)), "\n") {
		min := map[string]int{}

		for _, m := range diceRegex.FindAllStringSubmatch(s, -1) {
			n, _ := strconv.Atoi(m[1])
			min[m[2]] = slices.Max([]int{min[m[2]], n})
		}

		if min["red"] <= 12 && min["green"] <= 13 && min["blue"] <= 14 {
			part1 += i + 1
		}
		part2 += min["red"] * min["green"] * min["blue"]
	}
	return part1, part2
}
