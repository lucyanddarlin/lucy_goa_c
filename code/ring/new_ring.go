package main

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

func main() {

}
