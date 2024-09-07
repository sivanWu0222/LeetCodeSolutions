package main

import (
	"unicode"
)

/**
 * @Author: yirufeng
 * @Date: 2022/4/25 11:03 上午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/

/**
方法一：自己的解法
思路：使用一个map，key为字符，val为对应的行号
*/
//func findWords(words []string) []string {
//	hmap := map[byte]int{
//		'A': 1,
//		'B': 2,
//		'C': 2,
//		'D': 1,
//		'E': 0,
//		'F': 1,
//		'G': 1,
//		'H': 1,
//		'I': 0,
//		'J': 1,
//		'K': 1,
//		'L': 1,
//		'M': 2,
//		'N': 2,
//		'O': 0,
//		'P': 0,
//		'Q': 0,
//		'R': 0,
//		'S': 1,
//		'T': 0,
//		'U': 0,
//		'V': 2,
//		'W': 0,
//		'X': 2,
//		'Y': 0,
//		'Z': 2,
//	}
//
//	var flag bool
//
//	ret := []string{}
//
//	for _, word := range words {
//
//		start := hmap[byte(unicode.ToUpper(rune(word[0])))]
//		flag = true
//		for _, ch := range word {
//			if start != hmap[byte(unicode.ToUpper(ch))] {
//				flag = false
//				break
//			}
//		}
//
//		if flag {
//			ret = append(ret, word)
//		}
//	}
//	for _,v := range hmap {
//		fmt.Printf("%d", v)
//	}
//
//	return ret
//}

//方法二：使用数组进行改进
//使用一个数组进行改进，不需要使用hmap
func findWords(words []string) []string {
	//二十六个字母在的行号
	rowLines := "12210111011122000010020202"

	var flag bool

	ret := []string{}

	for _, word := range words {

		start := rowLines[unicode.ToUpper(rune(word[0]))-'A']
		flag = true
		for _, ch := range word {
			if start != rowLines[unicode.ToUpper(ch)-'A'] {
				flag = false
				break
			}
		}

		if flag {
			ret = append(ret, word)
		}
	}

	return ret
}
