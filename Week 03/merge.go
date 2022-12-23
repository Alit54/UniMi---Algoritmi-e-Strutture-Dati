/*
Author: Simone Alessandro Casciaro
Date: 23/12/2022
Lesson: 3.3
Text: Scrivete una funzione che implementa l’algoritmo ricorsivo mergeSort. La funzione:
	Scrivete una funzione che implementa l’algoritmo ricorsivo mergeSort. La funzione:
	1. divide il vettore in due sotto-vettori di dimensione circa uguale;
	2. ordina il sotto-vettore di sinistra richiamando se stessa;
	3. ordina il sotto-vettore di destra richiamando se stessa;
	4. integra (merge) i due vettori in un vettore ordinato.
	La base della ricorsione è, anche qui, data dai vattori di lunghezza 0 o 1, che sono sempre
	ordinati.
	La parte di integrazione (merge) di due vettore ordinati a1 e a2 funziona con un vettore di
	supporto: si scorrono entrambi i vettori da sinistra a destra usando due indicatori i1 e i2 rispettivamente; ad ogni passo si confronta a1[i1] con a2[i2] e si sceglie l’elemento più piccolo, lo si
	copia nel vettore di supporto (nella prima posizione libera) e si incrementa l’indicatore relativo
	ad esso. Quando i1 esce da a1 oppure i2 esce da a2, la parte rimanente dell’altro vettore viene
	copiata nel vettore di supporto. Alla fine si copia il contenuto del vettore di supporto nel vettore
	originale.*/

package main

func main() {
}

func MergeSort(slice []int) []int {
	out := make([]int, 0)
	// Base Case
	if len(slice) < 2 {
		return slice
	}
	// Recursive
	v1 := MergeSort(slice[:len(slice)/2])
	v2 := MergeSort(slice[len(slice)/2:])
	// Merge
	i, j := 0, 0
	for i < len(v1) && j < len(v2) {
		if v1[i] <= v2[j] {
			out = append(out, v1[i])
			i++
		} else {
			out = append(out, v2[j])
			j++
		}
	}
	if i < len(v1) {
		for _, v := range v1[i:] {
			out = append(out, v)
		}
	} else {
		for _, v := range v2[j:] {
			out = append(out, v)
		}
	}
	return out
}
