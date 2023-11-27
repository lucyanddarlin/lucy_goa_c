package main

import (
	"fmt"
	"sync"
)

type ArrayQueue struct {
	array []string
	size  int
	lock  sync.Mutex
}

// Push 入队
func (queue *ArrayQueue) Push(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.array = append(queue.array, v)

	queue.size++
}

// Pop 出队
func (queue *ArrayQueue) Pop() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("queue is empty")
	}

	v := queue.array[0]

	newArray := make([]string, queue.size-1)
	for i := 0; i < queue.size-1; i++ {
		newArray[i] = queue.array[queue.size-1-i]
	}

	queue.array = newArray
	queue.size--
	return v
}

func main() {
	queue := new(ArrayQueue)

	queue.Push("1")
	queue.Push("2")

	fmt.Println(queue.Pop())
	fmt.Println(queue.array)
}
