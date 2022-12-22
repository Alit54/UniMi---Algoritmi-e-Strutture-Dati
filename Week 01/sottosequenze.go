/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.13
	Text: Scrivere un programma che legge una sequenza di interi e stampa la somma degli interi in
		ciascuna sottosequenza crescente. Il programma interrompe la lettura quando riceve un segnale
		5 di EOF. Ad esempio, su input 1 2 13 0 7 8 9 -1 0 2 il programma stampa le somme 16, 24
		e 1. */

package main

import (
	"fmt"
	"io"
)

func main() {
	var prec, sum int
	number := 0
	fmt.Scan(&prec)
	sum += prec
	// Read input and print sums
	for {
		_, err := fmt.Scan(&number)
		if err == io.EOF {
			fmt.Println(sum)
			break
		}
		if number > prec {
			sum += number
		} else {
			fmt.Println(sum)
			sum = number
		}
		prec = number
	}
}
