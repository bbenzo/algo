package main

import "fmt"

type stack []int

type Stack interface {
	Len() int
	Pop() int
	Push(num int)
}

func NewStack() Stack {
	return &stack{}
}

func (s *stack) Len() int {
	return len(*s)
}

func (s *stack) Pop() int {
	var x int
	x, *s = (*s)[len(*s)-1], (*s)[:len(*s)-1]
	return x
}

func (s *stack) Push(num int) {
	*s = append(*s, num)
}

func moveOne(stack1, stack2 Stack) {
	item := stack1.Pop()
	if item == 0 {
		return
	}

	stack2.Push(item)
}


func moveTwo(start Stack, goal Stack, spare Stack) {
	moveOne(start, spare)
	moveOne(start, goal)
	moveOne(spare, goal)
}

func solve(first, second, third Stack, n int) {
	if n == 1 {
		moveOne(first, second)
		return
	} else if n == 2 {
		moveTwo(first, second, third)
		return
	} else {
		solve(first, third, second, n - 1)
		solve(first, second, third, 1)
		solve(third, second, first, n -1)
	}
}

func main() {

	/*
		RULES:
		1. You may only move one disk at a time -> move function
		2. No disk ever rest atop a smaller disk. Example: If disk 3 is on a stack, all disks below must have numbers > 3
	*/

	// setup
	stack1 := NewStack()
	stack2 := NewStack()
	stack3 := NewStack()

	stack1.Push(7)
	stack1.Push(6)
	stack1.Push(5)
	stack1.Push(4)
	stack1.Push(3)
	stack1.Push(2)
	stack1.Push(1)

	printStacks(stack1, stack2, stack3)

	// move recursively to stack2
	solve(stack1, stack2, stack3, stack1.Len())

	printStacks(stack1, stack2, stack3)
}

func printStacks(stack1 Stack, stack2 Stack, stack3 Stack) {
	fmt.Printf("\nStart: %v\n", stack1)
	fmt.Printf("Goal: %v\n", stack2)
	fmt.Printf("Spare: %v\n", stack3)
}
