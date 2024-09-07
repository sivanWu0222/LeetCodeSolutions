package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 9:54 上午
 * @Desc:
 **/

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, ok := hmap[target-nums[i]]; ok {
			return []int{index, i}
		}
		hmap[nums[i]] = i
	}

	return nil
}




