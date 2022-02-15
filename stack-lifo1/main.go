package main

import "fmt"

type stack []string

func main() {
	var s stack
	s.push("My")
	s.push("name")
	s.push("is")
	s.push("Ashutosh")

	fmt.Println(s)
	s.pop()
	fmt.Println(s)
}

func (s *stack) push(s1 string) {
	*s = append(*s, s1)
}

func (s *stack) pop() {
	if len(*s) == 0 {
		fmt.Println("Empty list")
		return
	}
	*s = (*s)[:len(*s)-1]
}
