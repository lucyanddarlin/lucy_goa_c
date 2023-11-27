package main

import "fmt"

type LinkNode struct {
	Data     int64
	NextNode *LinkNode
}

func main() {
	// 新的节点
	node := new(LinkNode)
	node.Data = 1

	// 新的节点 1
	node1 := new(LinkNode)
	node1.Data = 2
	node.NextNode = node1 // node1 链接到 node 上

	// 新的节点 2
	node2 := new(LinkNode)
	node2.Data = 2
	node1.NextNode = node2 // node2 链接到 node1 上

	// 新的节点 3
	node3 := new(LinkNode)
	node3.Data = 3
	node2.NextNode = node3 // node3 链接到 node2 上

	// 按顺序打印
	currentNode := node
	for {
		if currentNode != nil {
			fmt.Println(currentNode.Data)
			currentNode = currentNode.NextNode
			continue
		}
		break
	}
}
