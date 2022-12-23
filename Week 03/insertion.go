/*
	Author: Simone Alessandro Casciaro
	Date: 23/12/2022
	Lesson: 3.1
	Text: Scrivete una funzione che legga da standard input una sequenza di interi distinti terminati da
		0, memorizzandoli in un vettore ordinato (valutate se è più opportuno usare un array o una
		slice): ogni volta che viene letto un nuovo intero, il vettore viene scorso fino a trovare l’esatta
		collocazione del numero, quindi si crea lo spazio per il nuovo numero spostando in avanti i
		numeri successivi già memorizzati.
		Questo algoritmo è utile per riempire un vettore mantenendolo ordinato ad ogni passo. */

package main

import "fmt"

func main() {
	sort := make([]int, 0)
	bub := make([]int, 0)
	var num, i int
	// Insert first element in the slice
	fmt.Scan(&num)
	sort = append(sort, num)
	bub = append(bub, num)
	// Read and update slice, ordering it using InsertionSort
	for num != 0 {
		fmt.Scan(&num)
		bub = append(bub, num)
		for i = 0; i < len(sort); i++ {
			if num < sort[i] {
				slice2 := make([]int, 0)
				slice2 = append(slice2, sort[:i]...)
				slice2 = append(slice2, num)
				slice2 = append(slice2, sort[i:]...)
				sort = slice2
				break
			}
		}
		if i == len(sort) {
			sort = append(sort, num)
		}
	}
	fmt.Println(sort)
}
