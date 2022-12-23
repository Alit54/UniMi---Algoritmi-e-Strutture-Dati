/* Author: Simone Alessandro Casciaro
	Date: 12/12/2021 (Program was made by me when it went out)
	Lesson: 1.Extra
	Text: https://adventofcode.com/2021/day/6 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(partA(), partB())
}

func partA() uint64 {
	var n int
	var fish [9]uint64
	file, _ := os.Open("Advent.txt")
	scanner := bufio.NewScanner(file)
	// Memorise the lanternfishes
	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")
	for _, v := range input {
		n, _ = strconv.Atoi(v)
		fish[n]++
	}
	// Day cycle
	for day := 1; day <= 80; day++ {
		fish[0], fish[1], fish[2], fish[3], fish[4], fish[5], fish[6], fish[7], fish[8] = fish[1], fish[2], fish[3], fish[4], fish[5], fish[6], fish[0]+fish[7], fish[8], fish[0]
	}
	return fish[0] + fish[1] + fish[2] + fish[3] + fish[4] + fish[5] + fish[6] + fish[7] + fish[8]
}

func partB() uint64 {
	var n int
	var fish [9]uint64
	file, _ := os.Open("Advent.txt")
	scanner := bufio.NewScanner(file)
	// Memorise the lanternfishes
	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")
	for _, v := range input {
		n, _ = strconv.Atoi(v)
		fish[n]++
	}
	// Day cycle
	for day := 1; day <= 256; day++ {
		fish[0], fish[1], fish[2], fish[3], fish[4], fish[5], fish[6], fish[7], fish[8] = fish[1], fish[2], fish[3], fish[4], fish[5], fish[6], fish[0]+fish[7], fish[8], fish[0]
	}
	return fish[0] + fish[1] + fish[2] + fish[3] + fish[4] + fish[5] + fish[6] + fish[7] + fish[8]
}
