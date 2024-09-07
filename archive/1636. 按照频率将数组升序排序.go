/**
 @Author       cvenwu
 @Datetime     2022/9/19 22:32
 @Description
**/
package main

import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) []int {
	// 使用map统计数组元素出现次数
	count := make(map[int]int)

	// 统计次数
	for _, num := range nums {
		// 如果不存在这个元素，就将其次数置为1
		if _, ok := count[num]; !ok {
			count[num] = 1
		} else {
			count[num]++
		}
	}

	// 对nums按照次数进行排序
	sort.Slice(nums, func(i, j int) bool {
		// 按照次数升序排序，如果次数一样，按照对应的值降序排列
		return count[nums[i]] < count[nums[j]] || (count[nums[j]] == count[nums[i]] && nums[i] > nums[j])
	})
	// 返回结果
	return nums
}

func main() {
	fmt.Println(frequencySort(frequencySort([]int{1, 1, 2, 2, 2, 3})))
	fmt.Println(frequencySort(frequencySort([]int{2, 3, 1, 3, 2})))
}
