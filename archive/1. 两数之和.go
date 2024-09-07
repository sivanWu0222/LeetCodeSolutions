package main

/**
 * @Author: yirufeng
 * @Date: 2021/10/15 10:19 上午
 * @Desc:
 **/

//
// twoSum
// @Description:
// @param nums:
// @param target:
// @return []int:
//
func twoSum(nums []int, target int) []int {
	// key是元素在nums中的下标，val是元素的值
	hmap := make(map[int]int)
	for index, num := range nums {
		// 如果 目标值 - 当前遍历到的元素 出现在map中，就可以返回结果了
		if i, ok := hmap[target-num]; ok {
			return []int{index, i}
		}
		hmap[num] = index
	}

	// 走到这里说明结果为空
	return nil
}
