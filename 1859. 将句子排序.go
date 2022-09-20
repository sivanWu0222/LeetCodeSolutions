/**
 @Author       cvenwu
 @Datetime     2022/9/20 22:21
 @Description
**/
package main

import (
	"fmt"
	"strings"
)

func sortSentence(s string) string {
	// 1. 使用空格进行分割
	word_list := strings.Split(s, " ")
	// 2. 分配结果，单词个数就是最终结果列表长度
	ret := make([]string, len(word_list))
	// 3. 遍历每一个结果，并放到最终位置上，注意去掉最后的数字
	for _, word := range word_list {
		index := word[len(word)-1]
		// 这里要注意下标是从0开始，但是我们单词列表里面是1开始
		ret[index-'0'-1] = word[:len(word)-1] // 注意去除最后的数字
	}
	// 4. 返回结果
	return strings.Join(ret, " ")
}

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
}
