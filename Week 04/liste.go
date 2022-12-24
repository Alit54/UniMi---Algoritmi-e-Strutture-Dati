/*
Author: Simone Alessandro Casciaro
Date: 24/12/2022
Lesson: 4.1
Text: Scrivete un programma con l’implementazione di una lista concatenata semplice, seguendo i

	lucidi presentati a lezione (https://lonati.di.unimi.it/algolab-go/22-23/materiale/settimana04/liste-lucidi.pdf). Definite i tipi listNode e linkedList e scrivete queste funzioni:
	• newNode, che crea un nuovo nodo di lista;
	• addNewNode, che inserisce un nuovo nodo in testa alla lista;
	• printList, che stampa una lista;
	• searchList, che cerca un elemento in una lista;
	• removeItem, che cancella un item da una lista.
	Per testare le vostre funzioni scrivete una funzione main che gestisca un insieme dinamico
	(che variano nel tempo) di interi. Il main deve leggere da standard input una sequenza di istruzioni secondo il formato nella tabella qui sotto, dove n è un intero. I vari elementi sulla riga
	sono separati da uno o più spazi. Quando una riga è letta, viene eseguita l’operazione associata;
	le operazioni di stampa sono effettuate sullo standard output, e ogni operazione deve iniziare su
	una nuova riga.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var list linkedList
	for {
		scanner.Scan()
		if scanner.Text() == "f" {
			break
		}
		input := strings.Split(scanner.Text(), " ")
		switch input[0] {
		case "+":
			{
				num, _ := strconv.Atoi(input[1])
				addNewNode(&list, num)
			}
		case "-":
			{
				num, _ := strconv.Atoi(input[1])
				deleteItem(&list, num)
			}
		case "?":
			{
				num, _ := strconv.Atoi(input[1])
				fmt.Println(searchList(list, num))
			}
		case "c":
			{
				fmt.Println(count(list))
			}
		case "p":
			{
				printList(list)
			}
		case "o":
			{
				printReverse(list)
			}
		case "d":
			{
				for count(list) > 0 {
					element := list.head.item
					deleteItem(&list, element)
				}
			}
		case "default":
		}
	}
}

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
}

func newNode(val int) *listNode {
	node := new(listNode)
	node.item = val
	return node
}

func addNewNode(l *linkedList, val int) {
	node := newNode(val)
	node.next = l.head
	l.head = node
}

func printList(l linkedList) {
	p := l.head
	for p != nil {
		fmt.Print(p.item, " ")
		p = p.next
	}
	fmt.Println()
}

func printReverse(l linkedList) {
	for i := count(l); i >= 0; i-- {
		p := l.head
		for counter := 1; counter <= i; counter++ {
			if counter == i {
				fmt.Print(p.item, " ")
			}
			p = p.next
		}
	}
	fmt.Println()
}

func searchList(l linkedList, val int) (bool, *listNode) {
	p := l.head
	for p != nil {
		if p.item == val {
			return true, p
		}
		p = p.next
	}
	return false, nil
}

func deleteItem(l *linkedList, val int) bool {
	var curr, prev *listNode = l.head, nil
	for curr != nil {
		if curr.item == val {
			if prev == nil {
				l.head = l.head.next
			} else {
				prev.next = curr.next
			}
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

func count(l linkedList) int {
	counter := 0
	p := l.head
	for p != nil {
		counter++
		p = p.next
	}
	return counter
}
