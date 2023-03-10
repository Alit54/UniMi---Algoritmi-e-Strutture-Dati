package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type circNode struct {
	prec *circNode
	next *circNode
	val  int
}

type cerchio struct {
	head *circNode
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lista cerchio
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	aggiungiNodo(&lista, n, false)
	for scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
		aggiungiNodo(&lista, n, true)
	}
	fmt.Println(lista.head, lista.head.next, lista.head.next.next, lista.head.prec, lista.head.prec.prec)
}

func aggiungiNodo(list *cerchio, n int, flag bool) {
	node := &circNode{nil, nil, n}
	if !flag {
		node.next = node
		node.prec = node
	} else {
		node.next = list.head
		node.prec = list.head.prec
		node.prec.next = node
		node.next.prec = node
	}
	list.head = node
}
