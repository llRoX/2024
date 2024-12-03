package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func part2() int {

	path := "./input.txt"

	f, _ := os.ReadFile(path)
	data := string(f)

	re, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	dontRe, _ := regexp.Compile(`don't()|do()`)
	doOrs := dontRe.FindAllStringIndex(data, -1)

	matches := re.FindAllStringIndex(data, -1)
	sum := 0
	innerRe, _ := regexp.Compile(`\d+`)
	for _, match := range matches {
		start := match[0]
		apply := true
		for i, _ := range doOrs {
			v := doOrs[len(doOrs)-1-i]
			if v[0] < start {
				apply = v[1]-v[0] < 5
				break
			}
		}

		if !apply {
			continue
		}

		matchString := data[match[0]:match[1]]
		nMatches := innerRe.FindAllString(matchString, -1)
		n1, _ := strconv.Atoi(nMatches[0])
		n2, _ := strconv.Atoi(nMatches[1])
		sum += n1 * n2
	}

	return sum

}

func part1() int {
	path := "./input.txt"

	f, _ := os.ReadFile(path)
	data := string(f)

	re, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(data, -1)
	sum := 0
	for _, match := range matches {
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		sum += n1 * n2
	}

	return sum
}

func main() {
	fmt.Println(part1())
	//fmt.Println(part2())
}
