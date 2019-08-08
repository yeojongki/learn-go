package main

import (
	"fmt"
	"learn-go/queue"
)

func main() {
	q := queue.Queue{}
	q.Push(1)
	q.Push(2)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q, q.IsEmpty())
}
