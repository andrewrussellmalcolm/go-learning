package main

import (
	"fmt"
	"strings"
)

// Node :
type Node struct {
	data interface{}
	next *Node
	prev *Node
}

// String :
func (node Node) String() string {

	builder := strings.Builder{}
	builder.WriteString("node")

	return builder.String()
}

// LinkedList :
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// InsertAfter :
func (list *LinkedList) InsertAfter(node, newNode *Node) {
	newNode.prev = node
	if node.next == nil {
		list.tail = newNode

	} else {
		newNode.next = node.next
		node.next.prev = newNode
	}
	node.next = newNode
	list.size++
}

// InsertBefore :
func (list *LinkedList) InsertBefore(node, newNode *Node) {
	newNode.next = node
	if node.prev == nil {
		list.head = newNode
	} else {
		newNode.prev = node.prev
		node.prev.next = newNode
	}
	node.prev = newNode
	list.size++
}

// InsertHead :
func (list *LinkedList) InsertHead(newNode *Node) {
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		list.size++
		newNode.prev = nil
		newNode.next = nil
	} else {
		list.InsertBefore(list.head, newNode)
	}
}

// InsertTail :
func (list *LinkedList) InsertTail(newNode *Node) {
	if list.tail == nil {
		list.InsertHead(newNode)
	} else {
		list.InsertAfter(list.tail, newNode)
	}
}

// Remove :
func (list *LinkedList) Remove(node *Node) {
	if node.prev == nil {
		list.head = node.next
	} else {
		node.prev.next = node.next
	}

	if node.next == nil {
		list.tail = node.prev
	} else {
		node.next.prev = node.prev
	}
	list.size--
}

// Clear :
func (list *LinkedList) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

// Size :
func (list *LinkedList) Size() int {
	return list.size
}

// String :
func (list LinkedList) String() string {

	builder := strings.Builder{}

	builder.WriteString("\nlist\n")

	node := list.head
	for node != nil {
		builder.WriteString("\t")
		builder.WriteString(node.String())
		builder.WriteString("\n")
		node = node.next

		fmt.Println(node)
	}

	return builder.String()
}

// Main :
func main() {

	l := LinkedList{}

	a := Node{"A", nil, nil}
	l.InsertHead(&a)

	b := Node{"B", nil, nil}
	l.InsertAfter(&a, &b)

	c := Node{"C", nil, nil}
	l.InsertAfter(&b, &c)

	x := Node{"X", nil, nil}
	l.InsertHead(&x)

	p := Node{"P", nil, nil}
	l.InsertAfter(&b, &p)

	z := Node{"Z", nil, nil}
	l.InsertTail(&z)

	fmt.Println("Add 6 =====", l.Size(), l)

	l.Remove(&p)
	fmt.Println("Remove 1 =====", l.Size(), l)

	l.Clear()
	fmt.Println("Clear =====", l.Size(), l)

	for i := 0; i < 50; i++ {
		n := Node{fmt.Sprintf("entry %d", i), nil, nil}
		l.InsertTail(&n)
	}

	fmt.Println("Add 50 =====", l.Size(), l)

}
