package main

/**
 * @Author: yirufeng
 * @Date: 2022/4/25 4:04 下午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/

// 方法一：
//注意：数组的最开始k个元素最大的时候这种情况一定要考虑进去
//func findMaxAverage(nums []int, k int) float64 {
//	//最后的结果
//	sum := 0
//	ret := math.MinInt32 //不可以写0，因为可能是负数，
//
//	for i := 0; i < len(nums) - k + 1; i++ {
//
//		if i == 0 {
//			//计算前k个数字的和
//			for i := 0; i < k; i++ {
//				sum += nums[i]
//			}
//		} else {
//			//当前的范围是[i,i+k-1]
//			//需要在之前的基础上左边减去nums[i-1]，右边加上nums[i+k-1]，然后动态更新最大值
//			sum = sum - nums[i-1] + nums[i+k-1]
//		}
//
//		//动态更新最大值
//		if ret < sum {
//			ret = sum
//		}
//	}
//
//	return float64(ret) / float64(k)
//}

//方法一的第二种写法
func findMaxAverage(nums []int, k int) float64 {

	//保存连续k个的和
	sum := 0
	//最后的结果
	ret := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	ret = sum

	for i := 1; i < len(nums)-k+1; i++ {

		//当前的范围是[i,i+k-1]
		//需要在之前的基础上左边减去nums[i-1]，右边加上nums[i+k-1]，然后动态更新最大值
		sum = sum - nums[i-1] + nums[i+k-1]

		//动态更新最大值
		if ret < sum {
			ret = sum
		}
	}

	return float64(ret) / float64(k)
}
