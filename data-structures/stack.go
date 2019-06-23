package main

import "fmt"

type Stack struct {
	elements []int
}

func (s Stack) Size() int {
	return len(s.elements)
}

func (s *Stack) Push(v int) {
	// this will be represented in reverse order in memory
	s.elements = append(s.elements, v)
}

func (s *Stack) Pop() (int, error) {
	size := s.Size()

	if size == 0 {
		return 0, fmt.Errorf("index out of boundaries")
	}

	i, elements := s.elements[size-1], append(s.elements[:size-1])
	s.elements = elements

	return i, nil
}

func main() {
	s := Stack{}
	fmt.Printf("Size: %v\n", s.Size())
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("Size: %v\n", s.Size())
	v, _ := s.Pop()
	fmt.Printf("Value: %v\n", v)
	fmt.Printf("Size: %v\n", s.Size())

}
