/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.4
	Text: Scrivete le tre funzioni specificate sotto. Per testarle potete utilizzare il main riportato qui (con
		gli opportuni adattamenti).
		1. Scrivere una funzione stranoProdotto(numeri []int) int che, data come parametro
			una slice di interi, trovi quelli che sono maggiori di 7 e multipli di 4 e ne restituisca il
			prodotto. Ad esempio, se la slice contiene i numeri 12, 3, 4, 8, 9, 2, la funzione dovrà
			restituire il numero 96 (pari al prodotto di 12 per 8).
		2. Scrivere una funzione pariDispari(numeri []int) che, data come parametro una slice
			di interi, stampi, per ciascun numero, se è pari o dispari.
		3. Scrivere una funzione minDispari(numeri []int) int che, data una slice di interi,
			restituisca il più piccolo numero dispari (la slice può contenere sia numeri positivi che
			negativi).*/

package main

import (
	"fmt"
	"math"
)

// Main created by the prof for testing the 3 functions
func main() {
	const N = 10
	numeri := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}

	fmt.Println(stranoProdotto(numeri))
	pariDispari(numeri)
	fmt.Println(minDispari(numeri))
}

// Returns the strange product as defined in request 1
func stranoProdotto(values []int) int {
	product := 1
	for i := 0; i < len(values); i++ {
		if values[i] > 7 && values[i]%4 == 0 {
			product *= values[i]
		}
	}
	return product
}

// Prints if the numbers are odd or even as defined in request 2
func pariDispari(values []int) {
	for i := 0; i < len(values); i++ {
		if values[i]%2 == 0 {
			fmt.Println(values[i], "Even")
		} else {
			fmt.Println(values[i], "Odd")
		}
	}
}

// Returns the least odd number as defined in request 3
func minDispari(values []int) int {
	min := math.MaxInt
	for i := 0; i < len(values); i++ {
		if values[i]%2 == 1 && values[i] < min {
			min = values[i]
		}
	}
	return min
}
