/**
 @Author       cvenwu
 @Datetime     2022/9/10 22:44
 @Description
**/
package main

import "fmt"

//
// largestAltitude
// @Description: 使用两个变量，一个是结果，一个动态更新当前的海拔高度，一旦比结果大就更新结果
// @param gain:
// @return int:
//
func largestAltitude(gain []int) int {
	ret, cur := 0, 0
	for i := 0; i < len(gain); i++ {
		if cur+gain[i] > ret {
			ret = cur + gain[i]
		}
		cur += gain[i]
	}
	return ret
}

func main() {
	fmt.Println("123")
}
