/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.1
	Text: Scrivere un programma supera100.go che legge da standard input una sequenza di interi positivi terminata da -1 e stampa il primo numero che supera 100,
		se presente; altrimenti stampa “nessun numero maggiore di 100”. */

package main

import "fmt"

func main() {
	var number, i int
	slice := make([]int, 0)
	// Read input. Stop if you get -1
	for number != -1 {
		fmt.Scan(&number)
		slice = append(slice, number)
	}
	// Search number >100
	for i = 0; i < len(slice); i++ {
		if slice[i] > 100 {
			fmt.Println(slice[i])
			break
		}
	}
	// Print "nessun numero maggiore di 100" if no number is find
	if i == len(slice) {
		fmt.Println("nessun numero maggiore di 100")
	}
}
