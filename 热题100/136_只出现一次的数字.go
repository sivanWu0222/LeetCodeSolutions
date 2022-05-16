package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 4:30 下午
 * @Desc:
 **/
func singleNumber(nums []int) int {
	ret := 0

	for i := 0; i < len(nums); i++ {
		ret ^= nums[i]
	}

	return ret
}