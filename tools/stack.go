package tools

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int //栈大小
	Top    int //栈顶
	Array  [5]int
}

func (s *Stack) Push(val int) (err error) { //入栈
	if s.Top == s.MaxTop-1 { //判断栈是否满了
		return errors.New("stack full")
	}
	s.Top++
	s.Array[s.Top] = val
	return
}

func (s *Stack) List() { //显示栈
	if s.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("%d-->%d\n", i, s.Array[i])
	}
}

func (s *Stack) Pop() (val int, err error) { //出栈
	if s.Top == -1 { //判断栈是否为空
		return val, errors.New("stack empty")
	}
	val = s.Array[s.Top]
	s.Top--
	return val, nil
}
