package main

import "fmt"

type stack []string

func main() {
	var s stack
	s.push("My")
	s.push("name")
	s.push("is")
	s.push("Ankita")

	fmt.Println(s)
	result := s.pop()
	fmt.Println(result)
	fmt.Println(s)
}

func (s *stack) push(s1 string) {
	*s = append(*s, s1)
}

func (s *stack) pop() string {
	if len(*s) == 0 {
		fmt.Println("Empty list")
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
