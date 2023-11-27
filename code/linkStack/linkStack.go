package main

import (
	"fmt"
	"sync"
)

type LinkArray struct {
	topNode *LinkNode
	size    int
	lock    sync.Mutex
}

type LinkNode struct {
	Next  *LinkNode
	Value string
}

// Push 入栈
func (stack *LinkArray) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.topNode == nil {
		// 空栈, 初始化
		stack.topNode = &LinkNode{Value: v}
	} else {
		// 原来的链表
		preNode := stack.topNode

		// 新的节点
		newNode := &LinkNode{Value: v}

		newNode.Next = preNode

		stack.topNode = newNode
	}

	stack.size++
}

// Pop 出栈
func (stack *LinkArray) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.IsEmpty() {
		panic("stack is empty")
	}

	topNode := stack.topNode
	stack.topNode = topNode.Next

	v := topNode.Value

	stack.size--
	return v
}

// IsEmpty 栈是否为空
func (stack *LinkArray) IsEmpty() bool {
	return stack.Size() == 0
}

// Size 获取栈大小
func (stack *LinkArray) Size() int {
	return stack.size
}

// ToSlice 将链表转换为 slice
func (stack *LinkArray) ToSlice() []string {
	nextSlice := make([]string, stack.Size())

	if stack.IsEmpty() {
		return nextSlice
	}

	node := stack.topNode
	for {
		nextSlice = append(nextSlice, node.Value)
		node = node.Next
		if node == nil {
			break
		}
	}

	return nextSlice
}

func main() {
	stack := new(LinkArray)

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	fmt.Println(stack.Pop())
	fmt.Println(stack.ToSlice())
}
