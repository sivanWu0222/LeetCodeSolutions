package main

/**
 * @Author: yirufeng
 * @Date: 2021/10/15 10:19 上午
 * @Desc:
 **/

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		//首先检查target-nums[i]是否在hmap中
		if index, ok := hmap[target-nums[i]]; ok {
			return []int{index, i}
		}
		//加入到map中
		hmap[nums[i]] = i
	}

	//走到这里，说明最后结果是不存在的
	return nil
}

