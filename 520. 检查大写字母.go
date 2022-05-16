package main

import "unicode"

/**
 * @Author: yirufeng
 * @Date: 2022/4/24 12:31 下午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/


/**
方法一：思路，检查大小写的抽取出一个函数
 */
func detectCapitalUse(word string) bool {
	//如果长度为0或1直接返回true
	if len(word) <= 1 {
		return true
	}

	//如果第1个是小写，则全部都要小写
	if word[0] <= 'z' && word[0] >= 'a' {
		return checkCase(word, 0, false)
	} else if (word[0] <= 'Z' && word[0] >= 'A') && (word[1] <= 'Z' && word[1] >= 'A') { //如果第1个是大写，且第2个也是大写，则全都大写
		return checkCase(word, 2, true)
	} else if (word[0] <= 'Z' && word[0] >= 'A') && (word[1] <= 'z' && word[1] >= 'a') { //如果第1个是大写，第2个是小写，后面的全部都要小写
		return checkCase(word, 2, false)
	} else {
		return false
	}
}

//检查字母全部是否都为大写或者小写
//如果flag为true检查大写
//如果flag为false检查小写
func checkCase(word string, i int, flag bool) bool {

	var l, r byte
	if flag {
		l, r = 'A', 'Z'
	} else {
		l, r = 'a', 'z'
	}

	for i < len(word) {
		if word[i] < l || word[i] > r {
			return false
		}
		i++
	}
	return true
}

/**
方法二：
思路：
1. 如果长度小于等于1直接返回true
2. 如果有两个字符以及以上且第1个字符是大写，从第3个字符往后所有字符的大小写都要与第2个字符的大小写相同。
还有一种情况我们得考虑：字符串的第1个字符是小写的，第2个字符是大写的，直接返回false，（前提：字符串长度大于等于2）
 */
func detectCapitalUse(word string) bool {
	//如果长度小于等于1直接返回true
	if len(word) <= 1 {
		return true
	}

	//特殊情况：如果第1个字符是小写，第2个字符是大写，直接返回false
	if unicode.IsLower(rune(word[0])) && unicode.IsUpper(rune(word[1])) {
		return false
	}

	i := 2
	for i < len(word) {
		//如果与字符串的第2个字符大小写不一样直接返回false
		if unicode.IsLower(rune(word[i])) != unicode.IsLower(rune(word[1])) {
			return false
		}
		i++
	}

	return true
}

