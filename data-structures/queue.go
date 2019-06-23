package main

import "fmt"

type Queue struct {
	elements []int
}

func (q Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) Enqueue(v int) {
	q.elements = append(q.elements, v)
}

func (q *Queue) Dequeue() (int, error) {
	size := q.Size()

	if size == 0 {
		return 0, fmt.Errorf("index out of boundaries")
	}

	i, elements := q.elements[0], append(q.elements[1:])
	q.elements = elements

	return i, nil
}

func main() {
	q := Queue{}
	fmt.Printf("Size: %v\n", q.Size())
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Printf("Size: %v\n", q.Size())
	v, _ := q.Dequeue()
	fmt.Printf("Value: %v\n", v)
	fmt.Printf("Size: %v\n", q.Size())

}
