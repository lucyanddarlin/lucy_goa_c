package main

import "sync"

type Array struct {
	array []int
	len   int
	cap   int
	lock  sync.Mutex
}

func Make(len, cap int) *Array {
	if len > cap {
		panic("len is larger than cap")
	}

	s := new(Array)

	array := make([]int, len, cap)

	s.array = array
	s.cap = cap
	s.len = 0

	return s
}

func (a *Array) Append(e int) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.len == a.cap {
		// 长度达到容量需要扩容
		newCap := 2 * a.cap

		if a.cap == 0 {
			newCap = 1
		}

		newArray := make([]int, newCap, newCap)
		for k, v := range a.array {
			newArray[k] = v
		}
		a.array = newArray
		a.cap = newCap
	}

	a.array[a.len] = e
	a.len++
}

func (a *Array) MulAppend(e ...int) {
	for k := range e {
		a.Append(k)
	}
}

func (a *Array) Index(idx int) int {
	if a.len == 0 || idx > a.len {
		panic("idx is over len")
	}

	return a.array[idx]
}

func (a *Array) Len() int {
	return a.len
}

func (a *Array) Cap() int {
	return a.cap
}

func main() {

}
