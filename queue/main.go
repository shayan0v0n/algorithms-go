package main

import "fmt"

type queue []string

// fist in
func (s *queue) enqueue(element string) {
	*s = append([]string{element}, *s...)
}

// first out
func (s *queue) dequeue() {
	*s = (*s)[:len(*s)-1]
}

func (s *queue) show() queue {
	return *s
}

func main() {
	newqueue := queue{}
	newqueue.enqueue("1 - first element")
	newqueue.enqueue("2 - second element")
	newqueue.enqueue("3 - third element")
	newqueue.dequeue()
	newqueue.dequeue()
	newqueue.enqueue("4 - fourth element")

	fmt.Println(newqueue.show())
}
