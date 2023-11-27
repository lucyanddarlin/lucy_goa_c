package main

import "fmt"

// 每个 Node 节点, 包含前向和后向两个指针和数据域
type Node struct {
	next  *Node
	prev  *Node
	Value interface{}
}

type LinkedList struct {
	root   Node
	length int
}

// New 返回一个双向循环链表
func New() *LinkedList {
	return new(LinkedList).Init()
}

// Init 链表初始化, 指向自己, 形成一个环
func (l *LinkedList) Init() *LinkedList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.length = 0
	return l
}

func (l *LinkedList) LazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// Len 返回链表的长度
func (l *LinkedList) Len() int {
	return l.length
}

// insert 在指定节点后面插入节点
func (l *LinkedList) insert(at, e *Node) {
	// 1. 修改 e 的 prev
	e.prev = at
	// 2. 修改 e 的 next
	e.next = at.next
	// 3. 修改 e.next.prev
	e.next.prev = e
	// 修改 at 的 next
	at.next = e

	l.length++
}

// remove 删除特定节点
func (l *LinkedList) remove(e *Node) {
	// 1. 修改 n 节点的前一节点的 next 指针
	e.prev.next = e.next
	// 2. 修改 n 节点的后一节点的 prev 指针
	e.next.prev = e.prev

	e.next = nil
	e.prev = nil

	l.length--
}

// RangeSafe 从前向后遍历
//
//	callback 返回 true 就停止遍历
func (l *LinkedList) RangeSafe(callback func(e *Node) (exit bool)) {
	pos := l.root.next
	n := pos.next

	for pos != &l.root {
		if callback(pos) {
			break
		}

		pos = n
		n = pos.next
	}
}

// BackRangeSafe 从后向前遍历
//
//	callback 返回 true 就停止遍历
func (l *LinkedList) BackRangeSafe(callback func(e *Node) (exit bool)) {
	pos := l.root.prev
	n := pos.prev

	for pos != &l.root {
		if callback(pos) {
			break
		}

		pos = n
		n = pos.prev
	}
}

// Clear 清空链表
func (l *LinkedList) Clear() {
	l.RangeSafe(func(e *Node) (exit bool) {
		l.remove(e)
		return false
	})
}

// ToSlice 将链表转化为 slice
func (l *LinkedList) ToSlice() []any {
	if l.length == 0 {
		return nil
	}

	rv := make([]any, 0, l.length)
	for pos := l.root.next; pos != &l.root; pos = pos.next {
		rv = append(rv, pos.Value)
	}

	return rv
}

// testRangeSafe 测试正向遍历
func testRangeSafe(l *LinkedList) {
	fmt.Println("--------")
	l.RangeSafe(func(e *Node) (exit bool) {
		fmt.Println(e.Value)
		return false
	})
}

// testBackRangeSafe 测试反向遍历
func testBackRangeSafe(l *LinkedList) {
	fmt.Println("--------")
	l.BackRangeSafe(func(e *Node) (exit bool) {
		fmt.Println(e.Value)
		return false
	})
}

// testInsert 测试插入
func testInsert(l *LinkedList) {
	e1 := &Node{Value: 1}
	e2 := &Node{Value: 2}
	l.insert(&l.root, e1)
	l.insert(e1, e2)
}

// testClear 测试清空
func testClear(l *LinkedList) {
	l.Clear()
}

// testToSlice 测试 Slice 转换
func testToSlice(l *LinkedList) {
	fmt.Println("to slice", l.ToSlice())

}

func main() {
	l := new(LinkedList)
	l.Init()

	testInsert(l)
	testRangeSafe(l)
	testBackRangeSafe(l)
	testToSlice(l)
	testClear(l)

	fmt.Println("Link l's length is: ", l.Len())
}
