package 热题100

import "sort"

/**
 * @Author: yirufeng
 * @Date: 2021/10/20 9:23 上午
 * @Desc:
 **/

func merge(intervals [][]int) [][]int {
	var ret [][]int
	//1. 对区间按照起点进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//2. 向ret加入排序后的第一个元素
	ret = append(ret, intervals[0])

	//3. 遍历剩余的区间
	for i := 1; i < len(intervals); i++ {
		temp := ret[len(ret)-1][1]

		//如果当前遍历到的区间的起点大于结果中最后一个区间的终点，直接加入
		if temp < intervals[i][0] {
			ret = append(ret, intervals[i])
			continue
		}

		//走到这里说明，当前遍历的区间的起点小于结果中最后一个区间的终点
		//并且如果当前遍历到的区间的终点大于结果中最后一个区间的起点，则更新结果中最后一个区间的终点为较大值
		if temp < intervals[i][1] {
			ret[len(ret)-1][1] = intervals[i][1]
		}
	}

	//4. 返回结构
	return ret
}
