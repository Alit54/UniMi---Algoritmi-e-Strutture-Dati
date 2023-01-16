package main

import (
	"fmt"
)

type item struct {
	peso, valore int
}

func main() {
	types := []item{{5, 2}, {3, 3}, {7, 8}}
	fmt.Println(findValue(10, types...))
}

func findSemplValue(maxPeso int, items ...item) int {
	slice := make([]int, maxPeso+1)
	// Programmazione Dinamica
	for i := 0; i <= maxPeso; i++ {
		for t := 0; t < len(items); t++ {
			if i+items[t].peso <= maxPeso && slice[i+items[t].peso] < slice[i]+items[t].valore {
				slice[i+items[t].peso] = slice[i] + items[t].valore
			}
		}
	}
	return slice[maxPeso]
}

/*func findValue(maxPeso int, items ...item) int {
	row := make([]int, maxPeso+1)
	matrix := make([][]int, len(items))
	for i := 0; i < len(items); i++ {
		matrix[i] = row
	}
	for j := 0; j < len(matrix); j++ {
		for i := 0; i < len(matrix[j]); i++ {
			if j == len(items)-1 {
				if i+items[j].peso <= maxPeso && matrix[i+items[j].peso][j] < matrix[j][i]+items[j].valore {
					matrix[i+items[j].peso][j] = matrix[j][i] + items[j].valore
				}
			} else {
				if i+items[j].peso <= maxPeso && 0 < items[j].valore {
					matrix[j][i] = matrix[j][i] + items[j].valore
				}
			}
		}
	}
	return matrix[len(items)-1][maxPeso]
}*/

func findValue(maxPeso int, items ...item) int {
	row := make([]int, maxPeso+1)
	matrix := make([][]int, len(items))
	for i := 0; i < len(items); i++ {
		matrix[i] = row
	}
	fmt.Println(len(matrix), len(matrix[0]))
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {

		}
	}
	return 0
}
