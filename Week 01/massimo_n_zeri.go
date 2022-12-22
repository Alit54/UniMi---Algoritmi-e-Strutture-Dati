/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.12
	Text:  Scrivere un programma massimo_n_zeri.go che legge da standard input una sequenza di interi
		terminata da -1 e stampa il numero che contiene il maggior numero di 0. Nel caso in cui vi
		siano più numeri che contengono il massimo numero di 0, il programma stampa l?ultimo incontrato. Ad esempio, se la sequenza letta è 3040 145 80 1707 105002 78 1970 6005 -1
		il programma stampa 105002. */

package main

import "fmt"

func main() {
	number, max := 0, 0
	for number != -1 {
		fmt.Scan(&number)
		if howMany0(number) >= howMany0(max) {
			max = number
		}
	}
	fmt.Println(howMany0(max))
}

// Returns how many 0s has a number
func howMany0(num int) int {
	counter := 0
	for ; num != 0; num /= 10 {
		if num%10 == 0 {
			counter++
		}
	}
	return counter
}
