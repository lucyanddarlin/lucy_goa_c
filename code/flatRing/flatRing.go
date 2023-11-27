package main

import "fmt"

type Node struct {
	next, prev *Node
	value      interface{}
}

// Init 节点初始化
func (n *Node) Init() *Node {
	n.prev = n
	n.next = n

	return n
}

// Prev 获取上一个节点
func (n *Node) Prev() *Node {
	if n.prev == nil {
		return n.Init()
	}
	return n.prev
}

// Next 获取下一个节点
func (n *Node) Next() *Node {
	if n.next == nil {
		return n.Init()
	}
	return n.next
}

// Insert 添加节点
func (n *Node) Insert(e *Node) {
	e.prev = n
	e.next = n.next
	e.next.prev = e
	n.next = e
}

// Remove 删除指定节点
func (n *Node) Remove(e *Node) {
	e.next.prev = e.prev
	e.prev.next = e.next

	e.prev = nil
	e.next = nil
}

func main() {
	ns := &Node{value: 0}
	ns.Init()
	ns.Insert(&Node{value: 1})
	ns.Insert(&Node{value: 2})
	ns.Insert(&Node{value: 3})

	ns.Remove(ns.Next())

	pos := ns
	for {
		fmt.Println(pos.value)
		pos = pos.next

		if pos == ns {
			break
		}
	}
}
