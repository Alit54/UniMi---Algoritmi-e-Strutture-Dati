/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 2.2
	Text: Considerate la seguente funzione, che deve calcolare il valore massimo contenuto nel vettore
		numbers. La funzione max calcola il massimo tra i suoi due argomenti.
		1. Come deve essere completato il caso base?
			Assumete d’ora in poi che il vettore numbers contenga i valori
			[ 1, 2, 5, 7, -2, 10, 9, 21, 3, 8 ].
		2. Durante l’esecuzione della chiamata largest(numbers), considerate la chiamata ricorsiva che termina per prima. Qual è l’argomento passato in questa chiamata e quale valore
			restituisce la funzione?
		3. Con quali argomenti viene eseguita per la prima volta la funzione max
		4. E con quali argomenti viene eseguita l’ultima volta la funzione max? */

/*ANSWERS:
1. n == 1
2. largest(1). It'll return 1
3. max(1, 2)
4. max(8, 21)
*/

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 5, 7, -2, 10, 9, 21, 3, 8}
	fmt.Println(largest(numbers))
}

func largest(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return numbers[0]
	}
	return max(numbers[n-1], largest(numbers[:len(numbers)-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
