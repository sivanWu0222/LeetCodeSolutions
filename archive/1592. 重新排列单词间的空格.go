/**
 @Author       cvenwu
 @Datetime     2022/9/7 21:10
 @Description
**/
package main

import (
	"fmt"
	"strings"
)

// 给你一个字符串 text ，该字符串由若干被空格包围的单词组成。每个单词由一个或者多个小写英文字母组成，并且两个单词之间至少存在一个空格。题目测试用例保证 text 至少包含一个单词 。
//
// reorderSpaces
// @Description:
// @param text:
// @return string:
//
func reorderSpaces(text string) string {
	// 提前放置结果
	ret := make([]byte, len(text))
	// 1. 遍历一次字符串：统计空格个数 以及 单词个数
	backspaceCount := 0
	wordCount := 0
	// 保存单词列表
	var wordList []string
	for i := 0; i < len(text); {
		if text[i] == ' ' {
			i++
			backspaceCount++
			continue
		}

		// 走到这里说明不是空格，开始截取当前单词
		wordStart := i
		for i < len(text) && text[i] != ' ' {
			i++
		}
		wordList = append(wordList, text[wordStart:i])
		wordCount++
	}

	// 2. 每个单词后面放置 空格个数/(单词个数-1)
	index := 0
	for i := 0; i < len(wordList); i++ {
		// 首先放置单词
		for j := 0; index < len(ret) && j < len(wordList[i]); j++ {
			ret[index] = wordList[i][j]
			index++
		}
		// 避免 integer divide by zero
		if wordCount == 1 {
			break
		}
		// 放置空格
		for j := 0; index < len(ret) && j < backspaceCount/(wordCount-1); j++ {
			ret[index] = ' '
			index++
		}
	}

	// 3. 最后再往后填剩下的空格
	// 如果单词数为1， 避免 integer divide by zero
	if wordCount == 1 {
		for i := 0; i < backspaceCount; i++ {
			ret[len(ret)-1-i] = ' '
		}
	} else {
		for i := 0; i < backspaceCount%(wordCount-1); i++ {
			ret[len(ret)-1-i] = ' '
		}
	}

	return string(ret)
}

func main() {
	//s := "  this   is  a sentence "
	s := "  walks  udp package   into  bar a"
	sList := strings.Fields(s)
	strings.Count(s, " ")
	for i := 0; i < len(sList); i++ {
		fmt.Println(sList[i] + "!")
	}

	fmt.Printf("2%s2", reorderSpaces(s))
	fmt.Printf("2%s2", reorderSpaces("a"))
}
