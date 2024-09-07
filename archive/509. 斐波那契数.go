package main

import "fmt"

/**
 * @Author: yirufeng
 * @Date: 2022/4/24 3:29 下午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/

//方法一
//func fib(n int) int {
//
//	if n <= 1 {
//		return n
//	}
//
//	return fib(n-1) + fib(n-2)
//}

// 方法二：常数dp
func fib(n int) int {

	if n <= 1 {
		return n
	}

	prev, cur := 0, 1

	for i := 1; i < n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}

func main() {
	fmt.Println("123")
}
