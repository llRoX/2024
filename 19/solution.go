package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func canMatchPattern(rest string, options []string, cache map[string]bool) bool {
	if len(rest) == 0 {
		return true
	}

	v, exists := cache[rest]
	if exists && !v {
		return false
	}

	success := false
	for _, op := range options {
		l := len(op)
		if len(rest) < l || op != rest[:l] {
			continue
		}
		if len(rest) == l || canMatchPattern(rest[l:], options, cache) {
			success = true
			break
		}
	}

	cache[rest] = success

	return success
}

func canMatchPattern2(rest string, options []string, cache map[string]int) int {
	if len(rest) == 0 {
		return 0
	}

	v, exists := cache[rest]
	if exists {
		return v
	}

	ways := 0
	for _, op := range options {
		l := len(op)
		if len(rest) < l || op != rest[:l] {
			continue
		}
		if len(rest) == l {
			ways++
		} else {
			ways += canMatchPattern2(rest[l:], options, cache)
		}
	}

	cache[rest] = ways

	return ways
}

func part1() {
	path := "input.txt"
	f, _ := os.Open(path)

	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Scan()
	patternLine := sc.Text()
	patterns := strings.Split(patternLine, ", ")

	sc.Scan()
	count := 0
	cache := make(map[string]bool)
	for sc.Scan() {
		fullPattern := sc.Text()

		if canMatchPattern(fullPattern, patterns, cache) {
			count++
		}
	}
	fmt.Println(count)

}

func part2() {
	path := "input.txt"
	f, _ := os.Open(path)

	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Scan()
	patternLine := sc.Text()
	patterns := strings.Split(patternLine, ", ")

	sc.Scan()
	count := 0
	cache := make(map[string]int)
	for sc.Scan() {
		fullPattern := sc.Text()

		count += canMatchPattern2(fullPattern, patterns, cache)
	}
	fmt.Println(count)
}

func main() {
	part1()
	part2()
}
