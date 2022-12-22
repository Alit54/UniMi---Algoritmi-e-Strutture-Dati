/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.7
	Text: Si vuole scrivere un programma che legga una sequenza di stringhe e stampi, per ogni stringa, la
		sua prima vocale (per semplicità assumiamo che le stringhe contengano solo lettere minuscole).
		Nel caso in cui una stringa non contenga vocali il programma stampa “la parola non contiene
		vocali”. */

package main

import "fmt"

func main() {
	const N = 3
	strings := make([]string, N)
	// Read input
	for i := 0; i < N; i++ {
		fmt.Scan(&strings[i])
	}
	// Find first vocal
	for i := 0; i < len(strings); i++ {
		fmt.Println(firstVocal(strings[i]))
	}
}

// Returns the first vocal of a word. Returns "la parola non contiene vocali" if no vocal is found
func firstVocal(s string) string {
	for _, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			return string(v)
		}
	}
	return "la parola non contiene vocali"
}
