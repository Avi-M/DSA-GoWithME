package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) Size() int {
	return l.length
}
func (l *linkedList) isEmpty() bool {
	return l.length == 0
}
func (l linkedList) Peek() (int, bool) {
	if l.isEmpty() {
		fmt.Println("empty list exception")
		return 0, false
	}
	return l.head.data, true
}

func (l *linkedList) AddHead(val int) {
	l.head = &node{val, l.head}
	l.length++
}
func (l *linkedList) AddTail(val int) {
	curr := l.head
	newNode := &node{val, nil}
	l.length++

	if curr == nil {
		l.head = newNode
		return
	}
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode

}
func (l *linkedList) Print() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.data, " ")
		temp = temp.next
	}
	fmt.Println("")
}
func main() {
	ll := linkedList{}
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.AddTail(1)
	ll.AddTail(2)
	ll.AddTail(3)
	ll.Print()
	fmt.Println(ll.Size())
	fmt.Println(ll.isEmpty())
	fmt.Println(ll.Peek())
}
