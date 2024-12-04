package main

import (
	"bufio"
	"fmt"
	"os"
)

func part2() int {
	path := "./input.txt"
	file, _ := os.Open(path)

	defer file.Close()

	var wordSearch [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wordSearch = append(wordSearch, []rune(line))
	}

	count := 0

	for i := 1; i < len(wordSearch)-1; i++ {
		for j := 1; j < len(wordSearch[0])-1; j++ {
			// Find Middle Character
			if wordSearch[i][j] == 'A' {

				foundMatch := wordSearch[i-1][j-1] == 'M' &&
					wordSearch[i-1][j+1] == 'M' &&
					wordSearch[i+1][j+1] == 'S' &&
					wordSearch[i+1][j-1] == 'S'
				if foundMatch {
					count++
				}

				foundMatch = wordSearch[i-1][j-1] == 'S' &&
					wordSearch[i-1][j+1] == 'M' &&
					wordSearch[i+1][j+1] == 'M' &&
					wordSearch[i+1][j-1] == 'S'
				if foundMatch {
					count++
				}

				foundMatch = wordSearch[i+1][j+1] == 'S' &&
					wordSearch[i-1][j-1] == 'M' &&
					wordSearch[i-1][j+1] == 'S' &&
					wordSearch[i+1][j-1] == 'M'
				if foundMatch {
					count++
				}

				foundMatch = wordSearch[i-1][j-1] == 'S' &&
					wordSearch[i+1][j+1] == 'M' &&
					wordSearch[i-1][j+1] == 'S' &&
					wordSearch[i+1][j-1] == 'M'
				if foundMatch {
					count++
				}

			}
		}
	}

	return count
}

func isMatch(predicate func(int, int, int) bool, i int, j int) bool {
	for k := 0; k < 4; k++ {
		if predicate(i, j, k) {
			return false
		}
	}
	return true
}

func part1() int {
	path := "./input.txt"
	file, _ := os.Open(path)

	defer file.Close()

	var wordSearch [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wordSearch = append(wordSearch, []rune(line))
	}

	var funcSlice []func(int, int, int) bool
	wordToFind := "XMAS"

	funcSlice = append(funcSlice, func(i int, j int, k int) bool {
		return i+k >= len(wordSearch) || wordSearch[i+k][j] != rune(wordToFind[k])
	})

	funcSlice = append(funcSlice, func(i int, j int, k int) bool {
		return i-k < 0 || wordSearch[i-k][j] != rune(wordToFind[k])
	})

	funcSlice = append(funcSlice, func(i int, j int, k int) bool {
		return j+k >= len(wordSearch[0]) || wordSearch[i][j+k] != rune(wordToFind[k])
	})

	funcSlice = append(funcSlice, func(i int, j int, k int) bool {
		return j-k < 0 || i+k >= len(wordSearch) || wordSearch[i+k][j-k] != rune(wordToFind[k])
	})

	funcSlice = append(funcSlice, func(i int, j int, k int) bool {
		return i+k >= len(wordSearch) || j+k >= len(wordSearch[0]) || wordSearch[i+k][j+k] != rune(wordToFind[k])
	})

	funcSlice = append(funcSlice, func(i int, j int, k int) bool {
		return i-k < 0 || j-k < 0 || wordSearch[i-k][j-k] != rune(wordToFind[k])
	})

	funcSlice = append(funcSlice, func(i int, j int, k int) bool {
		return j-k < 0 || wordSearch[i][j-k] != rune(wordToFind[k])
	})

	funcSlice = append(funcSlice, func(i int, j int, k int) bool {
		return i-k < 0 || j+k >= len(wordSearch[0]) || wordSearch[i-k][j+k] != rune(wordToFind[k])
	})

	matches := 0

	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[0]); j++ {
			if wordSearch[i][j] == 'X' {
				for _, f := range funcSlice {
					if isMatch(f, i, j) {
						matches++
					}
				}
			}
		}
	}

	return matches
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
