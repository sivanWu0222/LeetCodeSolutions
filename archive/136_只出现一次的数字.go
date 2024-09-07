package main

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 4:30 下午
 * @Desc: 利用性质：一个数字与自己异或运算之后等于0
 **/
func singleNumber(nums []int) int {
	ret := 0

	for i := 0; i < len(nums); i++ {
		ret ^= nums[i]
	}

	return ret
}
