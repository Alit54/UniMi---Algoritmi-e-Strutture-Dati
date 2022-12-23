/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 2.1
	Text: Considerate la seguente funzione, che deve restituire il prodotto dei suoi due argomenti Cosa si
		deve scrivere, nel caso base, al posto di AAAAA e di BBBBB?

	func multiply(x, y int) int {
		if AAAAAA {
			return BBBBB
		} else {
			return x + multiply(x, y-1)
		}
	}
*/
package main

import "fmt"

func main() {
	fmt.Println(multiply(7, 6))
}

func multiply(x, y int) int {
	if y == 0 {
		return 0
	} else {
		return x + multiply(x, y-1)
	}
}
