package main

import (
	"fmt"
)

/**
 * @Author: yirufeng
 * @Date: 2021/10/24 8:09 下午
 * @Desc: 判断一个数是否为2的幂次方
 **/

func isPowerOfTwo(n int) bool {
	//注意特殊的case
	if n == 0 {
		return false

	}

	return n&(n-1) == 0
}

func main() {
	s := "23"

	for _, val := range s {
		fmt.Printf("%T",val)
	}
}