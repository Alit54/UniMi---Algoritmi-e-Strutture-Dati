/*
Author: Simone Alessandro Casciaro
Date: 23/12/2022
Lesson: 3.2
Text: Scrivete una funzione che riceva una v di interi e la ordini usando l’algoritmo SelectionSort:
	alla fine dell’esecuzione, la v originaria passata come argomento dovrà risultare ordinata. Può
	essere utile restituire la stessa v (ordinata), ad esempio per poterla passare come argomento
	ad altre funzioni, come in fmt.Println(selectionSort(v)).
*/
package main

func main() {
}

func selectionSort(v []int) []int {
	for k := 0; k < len(v)-1; k++ {
		m := k
		for j := k + 1; j < len(v); j++ {
			if v[j] < v[m] {
				m = j
			}
		}
		v[m], v[k] = v[k], v[m]
	}
	return v
}
