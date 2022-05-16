package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/21 4:16 下午
 * @Desc:
 **/

func majorityElement(nums []int) int {
	major, times := nums[0], 1
	for i := 1; i < len(nums); i++ {
		//如果次数此时为0，直接替换到major并让次数为1
		if times == 0 {
			major, times = nums[i], 1
			continue
		}

		//如果不等于我们的主要元素，次数减去1
		if nums[i] != major {
			times--
		} else { //否则次数++
			times++
		}
	}
	return major
}
