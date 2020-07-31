package main

import (
	"fmt"
)

//Node struct
type Node struct {
	key  int
	prev *Node
	next *Node
}

//List Struct
type List struct {
	head *Node
}

func main() {

	fmt.Println("Doubly linked list")

	list := List{}

	list.insert(1)
	list.insert(2)
	list.insert(3)
	list.insert(4)
	list.display()

	list.delete(3)
	list.display()

}

func (L *List) insert(key int) {

	newNode := &Node{
		key: key,
	}

	if L.head == nil {
		L.head = newNode
	} else {

		curNode := L.head

		for curNode.next != nil {
			curNode = curNode.next
		}

		curNode.next = newNode
		newNode.prev = curNode
	}
}

func (L *List) display() {

	curNode := L.head

	for curNode != nil {

		fmt.Print(curNode.key)
		fmt.Print("->")
		curNode = curNode.next

	}
	fmt.Println("")
}

func (L *List) delete(a int) {

	if L.head == nil {
		return
	}

	curNode := L.head
	var preNode *Node
	var nexNode *Node

	if curNode.key == a {
		L.head = L.head.next
		return
	}

	for curNode != nil {
		if curNode.key == a {
			break
		}

		curNode = curNode.next

	}
	if curNode == nil {
		return
	}

	if curNode.prev != nil {
		preNode = curNode.prev

	}
	if curNode.next != nil {
		nexNode = curNode.next
	}

	preNode.next = nexNode
	nexNode.prev = preNode

}
