# LeetCodeSolutions
LeetCode solutions


## 刷题笔记
1. 检查某个字符是否为大写`unicode.isUpper()`或者小写`unicode.isLower()`
2. 检查某个字符是否为英文字母`unicode.isLetter()`
3. 检查某个字符是否为数字`unicode.isNumber()`
4. 提取字符串中所有的单词列表(无论字符串首尾是否有空格，并且有多少空格)：`strings.Fields(s)`
5. 统计字符串s中指定内容c出现的次数：`strings.Count(s, c)`


## 一些待解答的问题
1. 统计一个字符串中单词的个数(空格可以出现在任意位置，并且空格可以出现任意多次)



统计单词个数，以及空格个数，以及单词列表
```go
func getWordCountAndList(text string) (wordCount, backspaceCount int, wordList []string) {
	// 1. 遍历一次字符串：统计空格个数 以及 单词个数
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
	return
}
```