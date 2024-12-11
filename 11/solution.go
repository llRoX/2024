package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type MapKey struct {
	Value int
	Steps int
}

func explore(value int, remainingSteps int, cache map[MapKey]int) int {
	if remainingSteps == 0 {
		return 1
	}

	key := MapKey{Value: value, Steps: remainingSteps}
	v, exists := cache[key]

	if exists {
		return v
	}

	if value == 0 {
		nV := explore(1, remainingSteps-1, cache)
		cache[key] = nV
		return nV
	}

	vAsString := strconv.Itoa(value)
	vLength := len(vAsString)

	if vLength%2 == 0 {
		v1, _ := strconv.Atoi(vAsString[:vLength/2])
		v2, _ := strconv.Atoi(vAsString[vLength/2:])

		nV := explore(v1, remainingSteps-1, cache) + explore(v2, remainingSteps-1, cache)
		cache[key] = nV

		return nV
	}

	nV := explore(value*2024, remainingSteps-1, cache)
	cache[key] = nV

	return nV

}

func splitStones(inputPath string, blinkTimes int) int {
	f, _ := os.ReadFile(inputPath)
	data := string(f)

	re, _ := regexp.Compile(`\d+`)
	matches := re.FindAllString(data, -1)

	values := make([]int, 0)

	for _, v := range matches {
		vvv, _ := strconv.Atoi(v)
		values = append(values, vvv)
	}

	cache := make(map[MapKey]int)

	sum := 0
	for _, v := range values {
		sum += explore(v, blinkTimes, cache)
	}

	return sum
}

func part1() int {
	path := "./input.txt"
	return splitStones(path, 25)
}

func part2() int {
	path := "./input.txt"
	return splitStones(path, 75)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
