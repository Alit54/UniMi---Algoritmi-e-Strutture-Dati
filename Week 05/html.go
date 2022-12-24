/*
	Author: Simone Alessandro Casciaro
	Date: 24/12/2022
	Lesson: 5.3
	Text: Semplificando un po’ le cose, in questo contesto chiamiamo documento html una sequenza di
		tag del tipo <a> (detti tag di apertura) oppure </a> (detti tag di chiusura, dove a è una qualsiasi
		stringa di caratteri alfabetici.
		Diciamo che un documento html è ben formato se i tag sono correttamente annidati, ovvero
		l’ordine dei tag soddisfa questi due criteri:
		• per ogni tag di apertura esiste uno e un solo un tag di chiusura corrispondente
		• se due tag di apertura compaiono in un determinato ordine, i corrispondenti tag di chiusura
		devono comparire nell’ordine opposto.
		Ad esempio, la sequenza <a> <b> </b> <c> <d> </d> </c> </a> costituisce un documento
		html ben formato. Al contrario, la sequenza <a> <b> </a> </c> non lo è perchè il tag <b> non
		è mai chiuso e il tag </c> non è mai aperto. Analogamente, la sequenza <a> <b> </a> </b> non
		è un documento html perchè il tag </a> viene chiuso prima del tag </b>.
		Scrivete un programma che legga una sequenza di tag e stabilisca se costituisce un documento
		html ben formato oppure no. Per farlo, potete usare una pila e fare riferimento a questo ciclo:
		leggi un tag t;
		se t è un tag di apertura
		inserisci t nella pila ;
		se t è un tag di chiusura
		estrai il tag t2 dalla cima della pila ;
		se t e t2 non corrispondono
		il documento non ben formato ;
		Il ciclo andrà ripetuto finché non ci sono più tag da leggere. Alla fine, affinché il documento sia
		ben formato, la pila deve essere vuota (pensate a degli esempi di input in cui restano tag nella
		pila alla fine dell’esecuzione).
		Nota: che tipo di item contiene la pila questa volta?
		Se il documento non è ben formato, il programma deve stampare un messaggio d’errore.
		Esempi:
		• con la stringa
		<a> <b> </b> </c>
		il programma stampa il messaggio
		errore in pos 4
		• con la stringa
		<a> <b> </b> <c> <d> </d>
		il programma stampa il messaggio
		sono rimasti aperti i seguenti tag: <a> <c> */

package main

import (
	"fmt"
	"strings"
)

func main() {
	isOK("<a> <b> </b> </c>")
	isOK("<a> <b> </b> <c> <d> </d>")
	isOK("<a> <b> </b> <c> <d> </d> </c> </a>")
}

func isOK(espressione string) bool {
	pila := make([]string, 0)
	input := strings.Split(espressione, " ")
	for i := 0; len(pila) > 0 && i < len(input) || i == 0; i++ {
		switch input[i][1] {
		case '/':
			{
				element := pila[len(pila)-1]
				pila = pila[:len(pila)-1]
				if element[1:] != input[i][2:] {
					fmt.Println("Errore in posizione", i+1)
					return false
				}
			}
		default:
			pila = append(pila, input[i])
		}
	}
	if len(pila) != 0 {
		fmt.Println("Sono rimasti aperti i tag", pila)
		return false
	}
	fmt.Println("Stringa OK")
	return true
}
