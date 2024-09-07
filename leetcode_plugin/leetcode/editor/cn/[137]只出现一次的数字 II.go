package main

//给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
//
// 你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,2,3,2]
//输出：3
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,1,0,1,99]
//输出：99
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
// nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
//
//
// Related Topics 位运算 数组 👍 1239 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
// 思路1：假设每一个数都出现3次。对nums中元素个数去重，然后都乘以3，减去题目中数组中所有数的和，差值就等于我们要找的那个数的两倍
// 思路2：利用位运算，依次确定每一个二进制位
// 思路3：利用真值表也就是电路技巧
func singleNumber(nums []int) int {

}

//leetcode submit region end(Prohibit modification and deletion)
