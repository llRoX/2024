package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type ZeroPos struct {
	Row int
	Col int
}

type ZerosAndGrid struct {
	Zeros []ZeroPos
	Grid  [][]int
}

func GetLineAsNumbers(pattern *regexp.Regexp, line string) []int {
	matches := pattern.FindAllString(line, -1)
	ret := make([]int, 0)

	for _, v := range matches {
		value, _ := strconv.Atoi(v)
		ret = append(ret, value)
	}

	return ret

}

func findZerosAndBuildGrid(sc *bufio.Scanner) ZerosAndGrid {
	row := 0
	zeros := make([]ZeroPos, 0)
	grid := make([][]int, 0)
	re, _ := regexp.Compile(`\d`)
	for sc.Scan() {
		line := sc.Text()
		values := GetLineAsNumbers(re, line)

		for cIdx, v := range values {
			if v != 0 {
				continue
			}
			zeros = append(zeros, ZeroPos{Row: row, Col: cIdx})

		}
		grid = append(grid, values)
		row++
	}

	return ZerosAndGrid{
		Grid:  grid,
		Zeros: zeros,
	}
}

func part2() int {
	path := "./input.txt"

	f, _ := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	defer f.Close()

	sc := bufio.NewScanner(f)
	zeroesAndGrid := findZerosAndBuildGrid(sc)
	sum := 0
	for _, v := range zeroesAndGrid.Zeros {
		sum += explore2(v.Row, v.Col, -1, zeroesAndGrid.Grid)
	}

	return sum
}

func explore2(row int, col int, last int, grid [][]int) int {
	rowCount := len(grid)
	colCount := len(grid[0])

	if row < 0 || row >= rowCount || col < 0 || col >= colCount {
		return 0
	}

	cValue := grid[row][col]
	if cValue-last != 1 {
		return 0
	}

	if last == 8 {
		return 1
	}

	return explore2(row+1, col, cValue, grid) +
		explore2(row-1, col, cValue, grid) +
		explore2(row, col+1, cValue, grid) + explore2(row, col-1, cValue, grid)
}

func explore(row int, col int, last int, grid [][]int, visited [][]bool) int {
	rowCount := len(grid)
	colCount := len(grid[0])

	if row < 0 || row >= rowCount || col < 0 || col >= colCount {
		return 0
	}

	cValue := grid[row][col]
	if cValue-last != 1 {
		return 0
	}

	if last == 8 {
		if visited[row][col] {
			return 0
		}
		visited[row][col] = true
		return 1
	}

	return explore(row+1, col, cValue, grid, visited) +
		explore(row-1, col, cValue, grid, visited) +
		explore(row, col+1, cValue, grid, visited) + explore(row, col-1, cValue, grid, visited)
}

func part1() int {
	path := "./input.txt"

	f, _ := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	defer f.Close()

	sc := bufio.NewScanner(f)
	zerosAndGrid := findZerosAndBuildGrid(sc)
	grid := zerosAndGrid.Grid
	sum := 0
	for _, v := range zerosAndGrid.Zeros {
		visited := make([][]bool, 0, len(grid))
		for range grid {
			visited = append(visited, make([]bool, len(grid[0])))
		}

		sum += explore(v.Row, v.Col, -1, grid, visited)
	}

	return sum
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
