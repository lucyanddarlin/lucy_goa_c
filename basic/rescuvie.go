package main

import "fmt"

// Rescuvie 阶乘 (递归)
func Rescuvie(r int64) int64 {
	if r == 0 {
		return 1
	}
	return r * Rescuvie(r-1)
}

// RescuvieTail 阶乘 (尾递归)
func RescuvieTail(n, a int64) int64 {
	if n == 1 {
		return a
	}
	return RescuvieTail(n-1, a*n)
}

// 斐波那契数列 (尾递归)
func Fibonacci(n, a1, a2 int) int {
	if n == 0 {
		return a1
	}
	return Fibonacci(n-1, a2, a1+a2)
}

// BinarySearch 二分查找
func BinarySearch(array []int, target, l, r int) int {
	if l > r {
		// 越界,返回 -1
		return -1
	}

	mid := (l + r) / 2
	middleNum := array[mid]

	if target == middleNum {
		// 匹配, 返回结果
		return mid
	} else if target < middleNum {
		// 大于中间值, 从右边开始查找
		return BinarySearch(array, target, l, mid-1)

	} else {
		// 小于中间值, 从左边开始查找
		return BinarySearch(array, target, mid+1, r)
	}
}

func main() {
	// fmt.Println(Rescuvie(5))
	// fmt.Println(RescuvieTail(5, 1))
	// fmt.Println(Fibonacci(1, 1, 1))
	// fmt.Println(Fibonacci(2, 1, 1))
	// fmt.Println(Fibonacci(3, 1, 1))
	// fmt.Println(Fibonacci(4, 1, 1))
	// fmt.Println(Fibonacci(5, 1, 1))
	arr := []int{1, 2, 3, 4}
	fmt.Println(BinarySearch(arr, 3, 0, len(arr)-1))
}
