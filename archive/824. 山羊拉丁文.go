package main

import (
	"fmt"
	"strings"
)

/**
 * @Author: yirufeng
 * @Date: 2022/4/21 11:29 上午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/
func toGoatLatin(sentence string) string {

	var ret []string
	var temp string
	var j int
	//1. 分割字符串
	senSlice := strings.Split(sentence, " ")
	//2. 遍历字符串切片
	for index, sen := range senSlice {
		//如果以元音字母(可能大写可能小写)开头，在单词后添加"ma"
		if sen[0] == 'a' || sen[0] == 'e' || sen[0] == 'i' || sen[0] == 'o' || sen[0] == 'u' ||
			sen[0] == 'A' || sen[0] == 'E' || sen[0] == 'I' || sen[0] == 'O' || sen[0] == 'U'{
			temp = sen+"ma"
		} else { //如果单词以辅音字母开头（即，非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"。
			temp = sen[1:] + string(sen[0]) + "ma"
		}

		//加上小写的a
		j = 0
		for j < index+1 {
			temp += "a"
			j++
		}

		//加入到结果中
		ret = append(ret, temp)
	}

	return strings.Join(ret," ")
}

func main() {

	for i := 3; i < 3; i++ {
		fmt.Println("123")
	}

}
