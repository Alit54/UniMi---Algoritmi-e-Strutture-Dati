/*
	Author: Simone Alessandro Casciaro
	Date: 22/12/2022
	Lesson: 1.2
	Text: Scrivere un programma che legge da riga di comando un intero che rappresenta il saldo di un conto corrente.
		Il programma legge poi da standard input una serie di numeri interi che rappresentano spese da addebitare sul conto e stampa il
		saldo finale. La lettura si interrompe quando il saldo è <=0. */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input
	saldo, err := strconv.Atoi(os.Args[1])
	var spesa int
	if err != nil {
		fmt.Println("Il saldo non è un numero.")
	}
	// Read payments until balance becomes not-positive
	for saldo > 0 {
		fmt.Scan(&spesa)
		saldo = saldo - spesa
	}
	// Print balance
	fmt.Println("Saldo finale:", saldo)
}
