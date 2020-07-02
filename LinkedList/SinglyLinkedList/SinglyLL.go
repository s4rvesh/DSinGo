package main

import (
	"errors"
	"fmt"
)

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	data int
	next *Node
}

func main() {

	fmt.Println("Singly Linked List program")

	list := List{}

	list.insert(5)
	list.insert(6)
	list.insert(7)
	list.insert(8)
	list.insert(9)
	list.insert(1)
	list.display()
	list.delete(7)
	fmt.Println(list.search(8))
	list.display()

}

func (L *List) insert(val int) {

	newnode := &Node{
		data: val,
		next: nil,
	}

	if L.head == nil {
		L.head = newnode

	} else {
		tempnode := L.head

		for tempnode.next != nil {
			tempnode = tempnode.next
		}
		tempnode.next = newnode
	}
}

func (L *List) delete(val int) {

	if L.head == nil {
		panic(errors.New("empty list"))
	}

	var prev *Node

	cur := L.head

	for cur.data != val {

		if cur.next == nil {
			panic(errors.New("not found"))
		}

		prev = cur
		cur = cur.next
	}

	prev.next = cur.next

}

func (L *List) display() {

	tnode := L.head

	for tnode != nil {
		fmt.Print(tnode.data)
		fmt.Print("->")
		tnode = tnode.next
	}
	fmt.Print("nil")
	fmt.Println(" ")
}

func (L *List) search(val int) int {

	tnode := L.head

	i := 0
	for tnode != nil {
		if tnode.data == val {
			return i
		}
		tnode = tnode.next
		i++
	}

	return 9999
}
