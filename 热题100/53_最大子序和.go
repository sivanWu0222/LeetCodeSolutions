package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/20 9:40 上午
 * @Desc:
 **/

func maxSubArray(nums []int) int {
	//全程都是我们要求的结果，需要不断更新
	ret := nums[0]

	//表示以第i个元素结尾时的连续子数组的最大和
	f := nums[0]

	for i := 1; i < len(nums); i++ {
		//如果前一个元素结尾的连续子数组的最大和是 >= 0 直接加上当前元素
		if f >= 0 {
			f += nums[i]
		} else {
			f = nums[i]
		}


		//动态更新我们的结果
		if ret < f {
			ret = f
		}
	}

	return ret
}
