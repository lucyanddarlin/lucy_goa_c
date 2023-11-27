package main

type Ring struct {
	value      interface{}
	prev, next *Ring // 前驱节点和后驱节点
}

// 初始化循环链表 O(1)
func (r *Ring) init() {
	r.next = r
	r.prev = r
}

// Create 创建指定大小的循环链表 O(n)
func Create(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}

	p.next = r
	r.prev = p

	return r
}

// func main() {
// 	r := new(Ring)
// 	r.init()
// }
