package main

import "fmt"

type stack []string

// fist in
func (s *stack) push(element string) {
	*s = append(*s, element)
}

// last out
func (s *stack) pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *stack) show() stack {
	return *s
}

func main() {
	newStack := stack{}
	newStack.push("1 - first element")
	newStack.push("2 - second element")
	newStack.push("3 - third element")
	newStack.pop()
	newStack.pop()
	newStack.push("4 - fourth element")

	fmt.Println(newStack.show())
}
