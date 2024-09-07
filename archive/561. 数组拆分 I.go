package main

import "sort"

/**
 * @Author: yirufeng
 * @Date: 2022/4/25 3:38 下午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/
//最后的结果最大，所以每一项都应该最大，但是每一项又都是两个数的最小值，所以要进行排序，然后从最大的开始相邻成对，此时每一项都是最大的了，我们结果也就最大了
//自己的想法：从小到大排序，然后从后往前两两成对
func arrayPairSum(nums []int) int {
	ret := 0
	//1. 排序
	sort.Ints(nums)
	//2. 从最后的一个元素开始，两两成对，计算结果并返回结果
	for i := len(nums) - 1; i >= 0; i -= 2 {
		ret += nums[i-1]
	}
	//3. 返回结果
	return ret
}
