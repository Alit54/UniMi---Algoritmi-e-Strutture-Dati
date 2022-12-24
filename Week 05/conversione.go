/*
	Author: Simone Alessandro Casciaro
	Date: 24/12/2022
	Lesson: 5.2
	Text: La pila è utile anche per convertire un’espressione da notazione infissa (con parentesi che racchiudono ogni operazione) a notazione postfissa. In questo caso, la pila non conterrà i numeri,
		ma gli operatori, rappresentati da caratteri.
		L’algoritmo è il seguente:
		leggi un token ;
		se il token è un numero
		stampa il token ;
		se il token è un operatore
		inserisci il token in cima alla pila ;
		se il token è una parentesi aperta
		ignora il token
		se il token è una parentesi chiusa
		estrai l’operatore in cima alla pila;
		stampalo;
		Notate che gli operandi devono apparire nello stesso ordine in entrambe le notazioni (la cosa è
		rilevante nel caso di operazioni non commutative, come la sottrazione); inoltre si può osservare
		che le parentesi aperte non sono necessarie nella notazione postfissa (lo sarebbero invece nel
		caso di operazioni tra più di due operandi).
		Scrivete una funzione func converti(espressione string) string che riceve una espressione in notazione infissa e restituisce l’espressione equivalente in notazione postfissa, usando
		l’algoritmo qui sopra (invece di stampare l’espressione convertita, dovrete costruire una stringa
		con l’espressione risultante).
		Ad esempio, ricevendo per argomento la stringa ( ( 5 - 3 ) * 2 ) , la funzione deve
		restituire la stringa 5 3 - 2 * .*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(converti("( ( 5 - 3 ) * 2 )"))
}

func converti(espressione string) string {
	out := ""
	pila := make([]string, 0)
	input := strings.Split(espressione, " ")
	for i := 0; i < len(input); i++ {
		_, err := strconv.Atoi(input[i])
		if err == nil {
			out += input[i] + " "
		} else {
			switch input[i] {
			case "+", "-", "*", "/":
				pila = append(pila, input[i]) // Push
			case ")":
				{
					op := pila[len(pila)-1]
					pila = pila[:len(pila)-1]
					out += op + " "
				}
			}
		}
	}
	return out
}
