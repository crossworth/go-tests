package main

import "fmt"

type Node struct {
	data int
	next *Node
	// we could have a prev, to make it a double linked list and facilitate the iteration of the list
}

type List struct {
	first *Node // We could have a tail and store the list in the "insert" order as well]
	// the advantage of a tail is that we can get the bottom of the list more easily
}

func (l *List) Size() int {
	size := 0

	for c := l.first; c != nil; c = c.next {
		size++
	}

	return size
}

func (l *List) InsertAfter(n Node, new Node) {
	for c := l.first; c != nil; c = c.next {
		if c.data == n.data {
			new.next = c.next
			c.next = &new
		}
	}
}

func (l *List) InsertFirst(new Node) {
	new.next = l.first
	l.first = &new
}

func (l *List) RemoveAfter(node Node) {
	for c := l.first; c != nil; c = c.next {
		if c.data == node.data && c.next != nil {
			if c.next.next != nil {
				c.next = c.next.next
			} else {
				c.next = nil
			}
		}
	}
}

func (l *List) RemoveFirst() {
	if l.first != nil {
		l.first = l.first.next
	}
}

func (l *List) Print() {
	for c := l.first; c != nil; c = c.next {
		fmt.Println(c.data)
	}
}

func main() {
	list := List{}
	//fmt.Printf("Size %v\n", list.Size())
	list.InsertFirst(Node{data: 10})
	list.InsertFirst(Node{data: 20})
	list.InsertFirst(Node{data: 30})
	list.Print()
	fmt.Printf("Size %v\n", list.Size())
	list.RemoveFirst()
	list.Print()
	fmt.Printf("Size %v\n", list.Size())
	list.InsertAfter(Node{data: 20}, Node{data: 55})
	list.Print()
	fmt.Printf("Size %v\n", list.Size())
	list.RemoveAfter(Node{data: 20})
	list.Print()
	fmt.Printf("Size %v\n", list.Size())
	//list.RemoveFirst()
	//fmt.Printf("Size %v\n", list.Size())
}
