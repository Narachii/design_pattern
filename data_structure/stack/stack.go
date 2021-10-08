package main

import (
	"fmt"
	"sync"
)

type customStack struct {
	stack []string
	lock sync.RWMutex
}

func (c *customStack) Push(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, name)
}

func (c *customStack) Pop() (string, error) {
	len := len(c.stack)
	if len > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.stack = c.stack[:len -1]
		val, _ := c.Front()
		return val, nil
	}
	return "", fmt.Errorf("Pop Error: Stack is empty")
}

func (c *customStack) Front() (string, error) {
	len := len(c.stack)
	if len > 0 {
		return c.stack[len-1], nil
	}
	return "", fmt.Errorf("Peep Error: Stack is empty")
}

func (c *customStack) Size() int {
	return len(c.stack)
}

func (c *customStack) Empty() bool {
	return len(c.stack) == 0
}

func main() {
	customStack := &customStack{
		stack: make([]string, 0),
	}
	fmt.Printf("Push: A\n")
	customStack.Push("A")
	fmt.Printf("Push: B\n")
	customStack.Push("B")
	for customStack.Size() > 0 {
		fmt.Printf("here")
		frontVal, _ := customStack.Pop()
		fmt.Printf("Pop: %s\n", frontVal)
	}
	fmt.Printf("Size: %d\n", customStack.Size())
}
