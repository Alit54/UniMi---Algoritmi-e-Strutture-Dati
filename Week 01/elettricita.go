/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.9
	Text: Vogliamo leggere una sequenza di N interi (almeno 3), che descrivono il consumo di elettricità
		in N giorni, e stampare i giorni in cui il consumo è stato inferiore rispetto sia al giorno prima
		che al giorno dopo. I giorni sono numerati a partire da 1. */

package main

import "fmt"

func main() {
	const N = 3
	numeri := make([]int, N)
	//Read input
	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}
	// Print days
	for i := 1; i < len(numeri)-1; i++ {
		if numeri[i] < numeri[i-1] && numeri[i] < numeri[i+1] {
			fmt.Println(i + 1)
		}
	}
}
