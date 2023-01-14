/*
	Author: Simone Alessandro Casciaro
	Date: 23/12/2022
	Lesson: 2.Extra
	Text: https://adventofcode.com/2021/day/9 (Part B) */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	row, column int
}

func main() {
	fmt.Println(partB())
}

func partB() int {
	var input [][]string
	var matrix [][]int
	var slice []int
	var num int
	var max [3]int
	var s string
	point := Position{0, 0}
	basin := make(map[Position]int) // Map of Low Points -> size of its basin
	file, _ := os.Open("Advent.txt")
	scanner := bufio.NewScanner(file)
	// save the inputs
	for scanner.Scan() {
		s = scanner.Text()
		input = append(input, strings.Split(s, ""))
	}
	// Convert input
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			num, _ = strconv.Atoi(input[i][j])
			slice = append(slice, num)
		}
		matrix = append(matrix, slice)
		slice = make([]int, 0)
	}
	// Calculate Low Points' Basins
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			point.row = i
			point.column = j
			if isLowPoint(point, matrix) {
				basin[point] = findBasin(point, matrix)
			}
		}
	}
	// Find 3 maximun values
	for _, v := range basin {
		if v > max[0] {
			max[2] = max[1]
			max[1] = max[0]
			max[0] = v
			continue
		}
		if v > max[1] {
			max[2] = max[1]
			max[1] = v
			continue
		}
		if v > max[2] {
			max[2] = v
		}
	}
	return max[0] * max[1] * max[2]
}

// Returns the size of a Low Point
func findBasin(point Position, matrix [][]int) int {
	if matrix[point.row][point.column] == 9 {
		return 0
	}
	counter := 1
	if point.row == 0 {
		// First row
		if point.column == 0 {
			// First column
			counter += findBasin(Position{point.row + 1, point.column}, matrix)
			counter += findBasin(Position{point.row, point.column + 1}, matrix)
		} else if point.column == len(matrix[point.row])-1 {
			// Last column
			counter += findBasin(Position{point.row + 1, point.column}, matrix)
			counter += findBasin(Position{point.row, point.column - 1}, matrix)
		} else {
			counter += findBasin(Position{point.row + 1, point.column}, matrix)
			counter += findBasin(Position{point.row, point.column - 1}, matrix)
			counter += findBasin(Position{point.row, point.column + 1}, matrix)
		}
	} else if point.row == len(matrix)-1 {
		// Last Row
		if point.column == 0 {
			// First column
			counter += findBasin(Position{point.row - 1, point.column}, matrix)
			counter += findBasin(Position{point.row, point.column + 1}, matrix)
		} else if point.column == len(matrix[point.row])-1 {
			// Last column
			counter += findBasin(Position{point.row - 1, point.column}, matrix)
			counter += findBasin(Position{point.row, point.column - 1}, matrix)
		} else {
			counter += findBasin(Position{point.row - 1, point.column}, matrix)
			counter += findBasin(Position{point.row, point.column - 1}, matrix)
			counter += findBasin(Position{point.row, point.column + 1}, matrix)
		}
	} else {
		// General case
		counter += findBasin(Position{point.row - 1, point.column}, matrix)
		counter += findBasin(Position{point.row + 1, point.column}, matrix)
		counter += findBasin(Position{point.row, point.column - 1}, matrix)
		counter += findBasin(Position{point.row, point.column + 1}, matrix)
	}
	return counter
}

// Returns if a Position is a Low Point
func isLowPoint(point Position, matrix [][]int) bool {
	if point.row == 0 {
		// First row
		if point.column == 0 {
			// First column
			if matrix[point.row][point.column] < matrix[point.row][point.column+1] && matrix[point.row][point.column] < matrix[point.row+1][point.column] {
				return true
			} else {
				return false
			}
		}
		if point.column == len(matrix[point.row])-1 {
			// Last column
			if matrix[point.row][point.column] < matrix[point.row][point.column-1] && matrix[point.row][point.column] < matrix[point.row+1][point.column] {
				return true
			} else {
				return false
			}
		}
		if matrix[point.row][point.column] < matrix[point.row][point.column+1] && matrix[point.row][point.column] < matrix[point.row][point.column-1] && matrix[point.row][point.column] < matrix[point.row+1][point.column] {
			return true
		} else {
			return false
		}
	}
	if point.row == len(matrix)-1 {
		// Last Row
		if point.column == 0 {
			// First column
			if matrix[point.row][point.column] < matrix[point.row][point.column+1] && matrix[point.row][point.column] < matrix[point.row-1][point.column] {
				return true
			} else {
				return false
			}
		}
		if point.column == len(matrix[point.row])-1 {
			// Last column
			if matrix[point.row][point.column] < matrix[point.row][point.column-1] && matrix[point.row][point.column] < matrix[point.row-1][point.column] {
				return true
			} else {
				return false
			}
		}
		if matrix[point.row][point.column] < matrix[point.row][point.column+1] && matrix[point.row][point.column] < matrix[point.row][point.column-1] && matrix[point.row][point.column] < matrix[point.row-1][point.column] {
			return true
		} else {
			return false
		}
	}
	// General case
	if matrix[point.row][point.column] < matrix[point.row][point.column+1] && matrix[point.row][point.column] < matrix[point.row][point.column-1] && matrix[point.row][point.column] < matrix[point.row-1][point.column] && matrix[point.row][point.column] < matrix[point.row+1][point.column] {
		return true
	} else {
		return false
	}
}
