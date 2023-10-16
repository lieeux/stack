package main

import (
	"fmt"
	"stack/tools"
)

func main() {
	stack := &tools.Stack{ //数栈
		MaxTop: 10,
		Top:    -1,
	}

	stack.Push(11)
	stack.Push(12)
	//stack.Push(13)
	//stack.Push(14)

	stack.List()

	val, _ := stack.Pop()
	val, _ = stack.Pop()
	fmt.Println(val)

	stack.List()
}
