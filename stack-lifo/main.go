package main

import (
	"errors"
	"fmt"
)

type stack []string

func main() {
	var s stack
	s.push("My")
	s.push("name")
	s.push("is")
	s.push("Ashutosh")

	fmt.Println(s)
	result, err := s.pop()
	if err != nil {
		fmt.Println("Empty list", err)
	}
	fmt.Println(result)
	fmt.Println(s)
}

func (s *stack) push(s1 string) {
	*s = append(*s, s1)
}

func (s *stack) pop() (string, error) {

	if len(*s) == 0 {
		return "", errors.New("Empty list")
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}
