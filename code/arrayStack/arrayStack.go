package main

import (
	"fmt"
	"sync"
)

type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

// Push 入栈
func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.array = append(stack.array, v)
	stack.size++
}

// Pop 出栈
func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("stack is empty")
	}

	// 栈顶元素
	top := stack.array[stack.size-1]

	// 创建新的数组, 空间占用不会越来越大, 但是可能移动次数比较多
	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}

	stack.array = newArray
	stack.size--

	return top
}

// Size 栈大小
func (stack *ArrayStack) Size() int {
	return stack.size
}

// IsEmpty 栈是否为空
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

func main() {
	stack := new(ArrayStack)
	stack.Push("1")
	stack.Push("2")
	fmt.Println(stack.array)
	fmt.Println(stack.Pop())
	fmt.Println(stack.array)
}
