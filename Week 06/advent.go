/*
	Author: Simone Alessandro Casciaro
	Date: 02/01/2023
	Lesson: 6.Extra
	Text: La sfida di Advent of Code, 2019, day 6, intitolata Universal Orbit Map, può essere affrontata
		modellando la situazione con un albero, e implementando l’albero con una tabella dei padri. La
		prima parte della sfida è riassunta qui sotto. Chi lo preferisce, può fare riferimento alla versione
		originale inglese, disponibile a questo indirizzo: https://adventofcode.com/2019/day/6
		Per entrambe le parti della sfida, prima di iniziare a scrivere codice, rispondete a queste
		domande:
			1. Modellate la situazione con un albero: cosa rappresentano i nodi dell’albero? cosa rappresenta la relazione padre/figlio?
			2. Riformulate i problemi usando la terminologia degli alberi:
				a) Parte 1 - Cosa sono le orbite dirette? E le orbite indirette? Come può essere descritto,
					in termini di alberi, il numero di orbite indirette di un oggetto?
				b) Parte 2 - Come può essere descritta, in termini di alberi, La distanza tra gli oggetti
					attorno a cui orbitano YOU e SAN?
			3. Progettate una soluzione per calcolare:
				a) Parte 1 - il numero di orbite indirette
				b) Parte 2 - il numero di trasferimenti di orbita necessari
					Ragionare in termini di alberi vi aiuterà a impostare gli algoritmi risolutivi!*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Son struct {
	name  string
	orbit int
}

func main() {
	answer := 0
	fathers := make(map[string]Son) // Son --> Father
	file, _ := os.Open("advent.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Create Map of Fathers
		row := strings.Split(scanner.Text(), ")")
		son := Son{row[0], 0}
		fathers[row[1]] = son
	}
	// Update
	
	// Calc
	for _, v := range fathers {
		answer += v.orbit
	}
	fmt.Println(answer)
}
