package main

import (
	"strconv"
)

/**
 * @Author: yirufeng
 * @Date: 2022/4/24 3:19 下午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/

//方法一：自己的解法
//func convertToBase7(num int) string {
//
//	if num == 0 {
//		return "0"
//	}
//
//	//使用一个变量保存正负
//	flag := true
//	if num < 0 {
//		flag = false
//		//转变成正数
//		num = -num
//	}
//
//	ret := []string{}
//	//得到正数的字符串
//	for num != 0 { //直到除数变为0
//		ret = append(ret, strconv.Itoa(num%7))
//		num /= 7
//	}
//
//	//反转ret数组
//	for i := 0; i < len(ret) >> 1; i++ {
//		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
//	}
//
//	if !flag {
//		//然后如果是负数前面加一个负号，否则直接输出
//		return "-" + strings.Join(ret,"")
//	}
//
//	return strings.Join(ret,"")
//
//}

//方法二：参考最快的解法
func convertToBase7(num int) string {

	sum := 0
	bit := 1

	for num != 0 {
		sum += (num % 7) * bit
		bit *= 10
		num /= 7
	}

	return strconv.Itoa(sum)
}
