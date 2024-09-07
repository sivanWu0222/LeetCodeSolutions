package main

/**
 * @Author: yirufeng
 * @Date: 2022/5/7 10:46 上午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:	双指针：一个指针指向下一个要放置的位置，另外一个指针指向下一个要放入的不重复元素所在数组的位置
 **/
func removeDuplicates(nums []int) int {

	if len(nums) <= 1 {
		return len(nums)
	}

	pos, ele := 1, 1

	for ele < len(nums) {
		//直到下一个要放入的元素与之前放置的最后一个元素不一样
		for ele < len(nums) && nums[ele] == nums[pos-1] {
			ele++
		}

		//Tips：这里要加上判断避免特殊case比如[1,1]会产生越界
		if ele < len(nums) {
			nums[pos], nums[ele] = nums[ele], nums[pos]
			pos, ele = pos+1, ele+1
		}

	}

	return pos
}