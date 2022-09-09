/**
 @Author       cvenwu
 @Datetime     2022/9/8 21:55
 @Description
**/
package main

import "fmt"

//
// constructArray
// @Description: 给你两个整数 n 和 k ，请你构造一个答案列表 answer ，该列表应当包含从 1 到 n 的 n 个不同正整数，并同时满足下述条件：
// 			假设该列表是 answer = [a1, a2, a3, ... , an] ，那么列表 [|a1 - a2|, |a2 - a3|, |a3 - a4|, ... , |an-1 - an|] 中应该有且仅有 k 个不同整数。
// 最关键的思路来自于https://leetcode.cn/problems/beautiful-arrangement-ii/solution/by-ac_oier-lyns/ 所说的间隔排列
// @param n: n个数，从1到n
// @param k: k个不同的差值
// @return []int:
//
func constructArray(n int, k int) []int {
	// 声明结果
	ret := make([]int, n)
	// 首先使用k个数字构造k-1个差值，间隔排列中，递增序列中保存为x, 递减排列保存为y, 然后继续添加后面的数字，使其保存固定大小的差值为1
	x, y := 1, n

	i := 0
	// 更新前k个数字
	for i < k {
		if i&1 == 0 { // 下标为偶数表示递增
			ret[i] = x
			x++
		} else { // 下标为奇数表示递减
			ret[i] = y
			y--
		}
		i++
	}

	// 此时x 一定 <= y
	// 填充剩下的数字
	// 更新剩下的n-k个数
	// 如果i此时为奇数，说明接下来从 x 到 y，也就是将未出现的数字从小到大
	// 如果i此时为偶数，说明接下来从 y 到 x，也就是将未出现的数字从大到小
	if i&1 == 0 {
		for y >= x {
			ret[i] = y
			y--
			i++
		}
	} else {
		for x <= y {
			ret[i] = x
			x++
			i++
		}
	}

	return ret
}

func main() {
	fmt.Println(constructArray(3, 1))
	fmt.Println(constructArray(3, 2))
	fmt.Println(constructArray(10, 7))
	//fmt.Println(constructArray(5, 2))
}
