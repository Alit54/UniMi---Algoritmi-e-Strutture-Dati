/*
	Author: Simone Alessandro Casciaro
	Date: 24/12/2022
	Lesson: 5.1
	Text: Per calcolare il valore di un espressione postfissa, si può usare un ciclo che esegue le seguenti
	azioni:
	leggi un token ( operatore o numero );
	se il token è un numero
	inseriscilo nella pila ;
	se il token è un operatore
	estrai due valori dalla pila ;
	applica ad essi l’operatore;
	inserisci il risultato nella pila;
	Alla fine la pila conterrà un solo valore, il risultato dell’espressone.
	Scrivete una funzione func valuta(espressione string) int che implementa il precedente algoritmo, utilizzando una pila. La funzione riceve una espressione in notazione postfissa
	e restituisce il suo valore.
	Ad esempio, se l’espressione è 5 3 - 2 *, il valore restituito deve essere 4.
	Per separare i token, si può usare la funzione strings.Split. Osservate che, per stabilire se
	token è un numero oppure un operatore, basta analizzare il suo primo carattere.
	Notate che in questo caso la pila dovrà contenere dei numeri. Per implementarla, potete usare
	una slice eseguendo le pop e le push nella posizione più a destra. In alternativa, potete usare una
	lista concatenata eseguendo le pop e le push in testa.
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(valuta("5 3 - 2 *"))
}

func valuta(espressione string) int {
	pila := make([]int, 0)
	input := strings.Split(espressione, " ")
	for i := 0; i < len(input); i++ {
		num, err := strconv.Atoi(input[i])
		if err == nil {
			// Push
			pila = append(pila, num)
		} else {
			// Double Pop
			elem := pila[len(pila)-2]
			ent := pila[len(pila)-1]
			// delete 2 elements
			pila = pila[:len(pila)-2]
			// Operation
			switch input[i] {
			case "+":
				pila = append(pila, elem+ent)
			case "*":
				pila = append(pila, elem*ent)
			case "/":
				pila = append(pila, elem/ent)
			case "-":
				pila = append(pila, elem-ent)
			}
		}
	}
	return pila[0]
}
