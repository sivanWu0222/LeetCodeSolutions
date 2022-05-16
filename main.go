package main

import (
	"fmt"
)

/**
 * @Author: yirufeng
 * @Date: 2021/8/23 4:33 下午
 * @Desc:
 **/

func main() {

	flag := true

	var l, r byte

	if flag {
		l, r = 'r', '1'
	}

	fmt.Println(l == 'r')
	fmt.Println(r == 'r')


	num := -10
	a := -num
	fmt.Println(a)


	s := []int{3,2,4}

	copy(s[1:], s[0:])

	s[0] = 100
	fmt.Println(s)

	fmt.Println((-3)%2)
}