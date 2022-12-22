/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.8
	Text: Si vuole scrivere un programma che legga una sequenza di dieci stringhe e stampi il numero di
		stringhe che cominciano con la lettera a e il numero di stringhe che cominciano con b. */

package main

import "fmt"

func main() {
	const N = 10
	countA, countB := 0, 0
	strings := make([]string, N)
	// Read input
	for i := 0; i < N; i++ {
		fmt.Scan(&strings[i])
	}
	// count letters
	for i := 0; i < len(strings); i++ {
		if firstLetter(strings[i]) == 'a' {
			countA++
		}
		if firstLetter(strings[i]) == 'b' {
			countB++
		}
	}
	// Prints Results
	fmt.Println(countA)
	fmt.Println(countB)
}

// Returns the first letter of a string
func firstLetter(s string) byte {
	return s[0]
}
