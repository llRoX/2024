package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func explore2(target int, current int, values []int) bool {
	if len(values) == 0 {
		return current == target
	}

	addTest := explore2(target, current+values[0], values[1:])
	if addTest {
		return true
	}
	multTest := explore2(target, current*values[0], values[1:])
	if multTest {
		return true
	}

	newValueForConcat := int(float64(current)*(math.Pow10(len(strconv.Itoa(values[0]))))) + values[0]
	concatTest := explore2(target, newValueForConcat, values[1:])

	return concatTest
}

func part2() int {
	path := "./input.txt"
	f, _ := os.Open(path)
	defer f.Close()
	sum := 0

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		text := sc.Text()
		match := strings.Split(text, ":")
		targetValue, _ := strconv.Atoi(match[0])
		splitValues := strings.Split(match[1][1:], " ")
		values := make([]int, 0)
		for _, val := range splitValues {
			cast, _ := strconv.Atoi(val)
			values = append(values, cast)
		}
		if explore2(targetValue, values[0], values[1:]) {
			sum += targetValue
		}
	}

	return sum

}

func explore(target int, current int, values []int) bool {
	if len(values) == 0 {
		return current == target
	}

	addTest := explore(target, current+values[0], values[1:])
	if addTest {
		return true
	}
	multTest := explore(target, current*values[0], values[1:])

	return multTest

}

func part1() int {
	path := "./input.txt"
	f, _ := os.Open(path)
	defer f.Close()
	sum := 0

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		text := sc.Text()
		match := strings.Split(text, ":")
		targetValue, _ := strconv.Atoi(match[0])
		splitValues := strings.Split(match[1][1:], " ")
		values := make([]int, 0)
		for _, val := range splitValues {
			cast, _ := strconv.Atoi(val)
			values = append(values, cast)
		}
		if explore(targetValue, values[0], values[1:]) {
			sum += targetValue
		}
	}

	return sum

}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
