package main

import (
	"errors"
	"fmt"
)

func CreateStack() []int {
	var statckTest []int
	return statckTest
}

func Push(stack []int, value int) []int {
	return append(stack, value)
}

func Pop(stack []int, value *int) ([]int, error) {
	if len(stack) == 0 {
		return stack, errors.New("stack is empty")
	}
	*value = stack[len(stack)-1]
	return stack[:len(stack)-1], nil
}
func Size(stack []int) int {
	return len(stack)
}

func TestStack() {
	stack := CreateStack()
	fmt.Printf("size = %d\n", Size(stack))
	stack = Push(stack, 4)
	stack = Push(stack, 5)
	stack = Push(stack, 6)
	stack = Push(stack, 7)
	fmt.Printf("size = %d\n", Size(stack))
	var value int
	var err error
	for i := 0; i < 5; i++ {
		stack, err = Pop(stack, &value)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Printf("pop = %d\n", value)
	}

}
