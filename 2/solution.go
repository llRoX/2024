package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func GetLineAsNumbers(pattern *regexp.Regexp, line string) []int {
	matches := pattern.FindAllString(line, -1)
	ret := make([]int, 0)

	for _, v := range matches {
		value, _ := strconv.Atoi(v)
		ret = append(ret, value)
	}

	return ret

}

func isSafe(values []int) bool {
	isDecreasing := values[0] > values[1]
	isSafe := true
	for i := 1; i < len(values); i++ {
		diff := (values[i] - values[i-1])
		if diff < 0 {
			diff = values[i-1] - values[i]
		}

		if diff < 1 || diff > 3 || (isDecreasing && values[i] > values[i-1]) || (!isDecreasing && values[i] < values[i-1]) {
			isSafe = false
			break
		}
	}

	return isSafe
}

func part2() int {
	path := "./input.txt"

	f, _ := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	defer f.Close()

	sc := bufio.NewScanner(f)
	re, _ := regexp.Compile(`\d+`)

	safeCount := 0
	for sc.Scan() {
		line := sc.Text()
		values := GetLineAsNumbers(re, line)

		if isSafe(values) {
			safeCount++
			continue
		}

		for i := 0; i < len(values); i++ {
			firstPart := values[:i]
			secondPart := values[i+1:]

			both := slices.Concat(nil, firstPart, secondPart)

			if isSafe(both) {
				safeCount++
				break
			}
		}

	}

	return safeCount
}

func part1() int {
	path := "./input.txt"

	f, _ := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	defer f.Close()

	sc := bufio.NewScanner(f)
	re, _ := regexp.Compile(`\d+`)

	safeCount := 0
	for sc.Scan() {
		line := sc.Text()
		values := GetLineAsNumbers(re, line)
		if isSafe(values) {
			safeCount++
		}

	}

	return safeCount
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
