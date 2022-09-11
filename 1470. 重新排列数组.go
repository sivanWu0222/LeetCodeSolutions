/**
 @Author       cvenwu
 @Datetime     2022/9/9 22:49
 @Description
**/
package main

import "fmt"

// 方法一：使用一个三个变量，第1个变量表示前半部分下一个应该放置的元素，第2个变量表示后半部分下一个应该放置的元素，第三个变量表示返回的结果中应该放置元素的位置
// shuffle
// @Description: 给你一个数组 nums ，数组中有 2n 个元素，按 [x1,x2,...,xn,y1,y2,...,yn] 的格式排列。请你将数组按 [x1,y1,x2,y2,...,xn,yn] 格式重新排列，返回重排后的数组。
// @param nums:
// @param n:
// @return []int:
//
func shuffle(nums []int, n int) []int {
	ret := make([]int, len(nums))
	length := len(nums)
	i, j := 0, length>>1
	for k := 0; k < length; {
		ret[k] = nums[i]
		k++
		i++
		ret[k] = nums[j]
		j++
		k++
	}

	return ret
}

// 按照LeetCode题解进一步优化

func shuffle(nums []int, n int) []int {
	ret := make([]int, n<<1)
	for i := 0; i < n; i++ {
		ret[2*i] = nums[i]
		ret[2*i+1] = nums[i+n]
	}
	return ret
}

func main() {
	fmt.Println(shuffle([]int{1, 2, 3, 4}, 2))
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
