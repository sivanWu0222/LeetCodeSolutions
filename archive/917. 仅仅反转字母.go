package main

import "unicode"

/**
 * @Author: yirufeng
 * @Date: 2022/4/25 7:44 下午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/

//思路：双指针进行对撞的思想，左边指针指向一个字母，右边指针指向一个字母
//func reverseOnlyLetters(s string) string {
//
//	//转换为byte数组
//	sTemp := []byte(s)
//
//	//l指向左侧字母，r指向右侧字母
//	l, r := 0, len(s)-1
//	for l < r {
//		//不断循环让l指向左侧字母
//		for l <= len(s)-1 && !unicode.IsLetter(rune(s[l])) {
//			l++
//		}
//
//		//不断循环让r指向右侧字母
//		//注意判断r没有越界
//		for r >= 0 && !unicode.IsLetter(rune(s[r])) {
//			r--
//		}
//
//		if l >= len(s) || r < 0 || l >= r {
//			break
//		}
//
//		//交换l与r
//		sTemp[l], sTemp[r] = sTemp[r], sTemp[l]
//		//然后改变l与r
//		l++
//		r--
//	}
//
//	return string(sTemp)
//}

// 对上面代码进行优化
func reverseOnlyLetters(s string) string {

	//转换为byte数组
	sTemp := []byte(s)

	//l指向左侧字母，r指向右侧字母
	l, r := 0, len(s)-1
	for l < r {
		//不断循环让l指向左侧字母
		for l < r && !unicode.IsLetter(rune(s[l])) {
			l++
		}

		//不断循环让r指向右侧字母
		//注意判断r没有越界
		for r > l && !unicode.IsLetter(rune(s[r])) {
			r--
		}

		if l >= r {
			break
		}

		//交换l与r
		sTemp[l], sTemp[r] = sTemp[r], sTemp[l]
		//然后改变l与r
		l++
		r--
	}

	return string(sTemp)
}
