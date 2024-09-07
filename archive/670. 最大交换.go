// Package main
// @Description:
package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	// 从开始向后一直找，直到找到一个不为最大的开始，然后与后面的最大的数字进行交换
	// 将整数转换为[]byte数组
	byteArr := []byte(strconv.Itoa(num))

	// 遍历byteArr
	// 假设当前遍历到的是
	n := len(byteArr)
	maxIndex := n - 1
	idx1, idx2 := -1, -1

	// 下面这个循环既可以做到边更新最大的值所在的下标，又可以做到找到比最大的值小的左侧下标
	for i := n - 1; i >= 0; i-- {
		if byteArr[i] > byteArr[maxIndex] {
			maxIndex = i
		} else if byteArr[i] < byteArr[maxIndex] {
			idx1, idx2 = i, maxIndex
		}
	}

	// 说明在前面没有可以与maxIndex交换的数
	if idx1 < 0 {
		return num
	}

	byteArr[idx1], byteArr[idx2] = byteArr[idx2], byteArr[idx1]
	// 转换为结果并输出
	ret, _ := strconv.Atoi(string(byteArr))
	return ret
}

func main() {
	fmt.Println(maximumSwap(7882319))
}
