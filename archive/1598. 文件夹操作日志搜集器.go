/**
 @Author       cvenwu
 @Datetime     2022/9/9 15:04
 @Description
**/
package main

import "fmt"

//
// minOperations
// @Description: 通过一个数来记录当前距离根目录有多远，如果拿到的是"./"不用修改，拿到的是"../"且距离大于0，让距离减去1，如果拿到的是其他的，则直接让距离+1即可，最后返回结果
// @param logs:
// @return int:
//
func minOperations(logs []string) int {
	// 声明最后返回的结果
	ret := 0

	// 遍历日志
	for _, log := range logs {
		// 如果是当前目录 或者是返回上一级目录，但是已经在根目录
		if log == "./" {
			continue
		}
		// 如果是上一级目录
		if log == "../" {
			if ret > 0 { // 如果是上一级目录，并且不在根目录
				ret--
			}
			continue
		}

		// 如果是其他目录
		ret++
	}

	return ret
}

func main() {
	logs := []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}
	fmt.Println(minOperations(logs))
	logs = []string{"d1/", "../", "../", "../"}
	fmt.Println(minOperations(logs))
	fmt.Println("123")
}
