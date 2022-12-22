/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.5
	Text: Scrivete le tre funzioni specificate sotto. Per testarle potete utilizzare il main riportato qui (con
		gli opportuni adattamenti).
		1. Scrivere una funzione quanteConA(parole []string) int che, data una slice di stringhe, restituisca quante stringhe contengono il carattere ‘a’.
		2. Scrivere una funzione primaConA(parole []string) string che, data una slice di
			stringhe, restituisca la prima parola che contentiene il carattere ‘a’, o la stringa vuota.
		3. Scrivere una funzione piuCorta(parole []string) int che, data una slice di stringhe,
			restituisca la lunghezza della stringa più corta in termini di byte.*/

package main

import "fmt"

// Main created by the prof for testing the 3 functions
func main() {
	const N = 10
	parole := make([]string, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&parole[i])
	}

	fmt.Println(quanteConA(parole))
	fmt.Println(primaConA(parole))
	fmt.Println(piuCorta(parole))
}

// Returns how many words contains 'a' as defined in request 1
func quanteConA(words []string) int {
	counter := 0
	for i := 0; i < len(words); i++ {
		for _, v := range words[i] {
			if v == 'a' {
				counter++
				break
			}
		}
	}
	return counter
}

// Returns the word required as defined in request 2
func primaConA(words []string) string {
	for i := 0; i < len(words); i++ {
		for _, v := range words[i] {
			if v == 'a' {
				return words[i]
			}
		}
	}
	// Word with 'a' not found
	return ""
}

// Returns the word with the less byte as defined in request 3
func piuCorta(words []string) int {
	min := len(words[0])
	for i := 1; i < len(words); i++ {
		if len(words[i]) < min {
			min = len(words[i])
		}
	}
	return min
}
