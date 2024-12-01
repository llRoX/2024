package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func part1() int {
	path := "./input.txt"

	f, _ := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	defer f.Close()

	sc := bufio.NewScanner(f)
	re, _ := regexp.Compile(`\d+`)
	distance := 0
	firsts := make([]int, 0)
	seconds := make([]int, 0)
	for sc.Scan() {
		line := sc.Text()
		matches := re.FindAllStringIndex(line, -1)

		left, _ := strconv.Atoi(line[matches[0][0]:matches[0][1]])
		right, _ := strconv.Atoi(line[matches[1][0]:matches[1][1]])
		firsts = append(firsts, left)
		seconds = append(seconds, right)

	}

	slices.Sort(firsts)
	slices.Sort(seconds)

	for idx, _ := range firsts {
		a := firsts[idx]
		b := seconds[idx]

		value := a - b
		if value < 0 {
			value = b - a
		}
		distance += value
	}

	return distance
}

func part2() int {
	path := "./input.txt"

	f, _ := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	defer f.Close()

	sc := bufio.NewScanner(f)
	re, _ := regexp.Compile(`\d+`)
	firsts := make([]int, 0)
	seconds := make(map[int]int)
	for sc.Scan() {
		line := sc.Text()
		matches := re.FindAllString(line, -1)

		left, _ := strconv.Atoi(matches[0])
		right, _ := strconv.Atoi(matches[1])

		secondsValue := seconds[right]

		firsts = append(firsts, left)
		seconds[right] = secondsValue + 1

	}
	score := 0
	for _, k := range firsts {
		b := seconds[k]
		score += k * b
	}

	return score
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
