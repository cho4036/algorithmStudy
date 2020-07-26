package main

import (
	"fmt"
	"math/big"
)

type node struct {
	value int
}

type stack struct {
	count int
	nodes []*node
}

func (s *stack) Push(n *node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *stack) Pop() *node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func NewStack() *stack {
	return &stack{}
}

func main() {
	n := 30
	fmt.Println(minSteps(n))
}

func minSteps(n int) int {
	stack := NewStack()
	sum := 0

	factorization(n, stack)

	for stack.count != 0 {
		sum += stack.Pop().value
	}

	return sum
}

func factorization(n int, s *stack) {
	if n == 1 {
		return
	}
	factor := find_factor(n)
	if factor == -1 {
		panic("Error")
	}
	s.Push(&node{value: factor})
	factorization(n/factor, s)

	return
}

func find_factor(n int) int {
	if n%2 == 0 {
		return 2
	}
	if isPrime(n) {
		return n
	}

	for i := 3; i*i <= n; i = i + 2 {
		if n%i == 0 {
			return i
		}
	}
	return -1
}

func isPrime(n int) bool {
	return big.NewInt(int64(n)).ProbablyPrime(0)
}
