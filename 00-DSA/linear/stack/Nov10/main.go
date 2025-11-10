package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	elements []int
}

func (s *Stack) isEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) push(value int) {
	s.elements = append(s.elements, value)
}

func (s *Stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("underflow error")
	}
	topI := len(s.elements) - 1
	topValue := s.elements[topI]
	s.elements = s.elements[:topI]
	return topValue, nil
}

func (s *Stack) peek() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("underflow error")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack) size() int {
	return len(s.elements)
}

func (s *Stack) display() {
	for _, v := range s.elements {
		fmt.Println("data is:", v)
	}
}

func main() {
	fmt.Println("Stack implementation")

	s := &Stack{}
	s.push(10)
	s.push(20)
	s.push(30)
	s.push(40)
	s.display()

	val, err := s.pop()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Popped:", val)
	}

	s.display()
}
