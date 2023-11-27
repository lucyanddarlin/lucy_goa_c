package main

import (
	"fmt"
	"sync"
)

type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Next  *LinkNode
	Value string
}

// Push 入队
func (queue *LinkQueue) Push(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.root == nil {
		// 空队列, 初始化
		queue.root = &LinkNode{Value: v}
	} else {
		// 添加至队列
		newNode := &LinkNode{Value: v}
		lastNode := queue.root
		for {
			if lastNode.Next == nil {
				break
			}
			lastNode = lastNode.Next
		}
		lastNode.Next = newNode
	}

	queue.size++
}

// Pop 出队
func (queue *LinkQueue) Pop() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("queue is empty")
	}

	v := queue.root.Value

	queue.root = queue.root.Next
	queue.size--

	return v
}

func (queue *LinkQueue) ToSlice() []string {
	nextSlice := make([]string, 0, queue.size)
	if queue.size == 0 {
		return nextSlice
	}

	node := queue.root

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
	queue := new(LinkQueue)

	queue.Push("1")
	queue.Push("3")
	queue.Push("4")

	fmt.Println(queue.ToSlice())
	fmt.Println(queue.Pop())
	fmt.Println(queue.ToSlice())

}
