package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	next := l.head
	l.head = n
	l.head.next = next
	l.length++
}

func (l *linkedList) printListData() {
	length := l.length
	current := l.head

	for length > 0 {
		fmt.Printf("%d ", current.data)
		current = current.next
		length--
	}
}

func (l *linkedList) deleteByValue(value int) {

}

func countNodes(n *node) int {
	count := 1
	current := n

	for current.next != nil {
		current = current.next
		count++
	}

	return count
}

func main() {
	node1 := node{data: 2}
	node2 := node{data: 5}
	node3 := node{data: 8}

	/* manually linkedNode
	node1.next = &node2
	node2.next = &node3
	fmt.Println(countNodes(&node1))
	fmt.Println(node1.next)
	*/

	// use prepend func to linkedNode
	l := linkedList{}
	l.prepend(&node3)
	l.prepend(&node2)
	l.prepend(&node1)
	l.printListData()
}
