package main

/**
 * @Author: yirufeng
 * @Date: 2022/5/11 12:29 PM
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/

func myPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}

	var ret float64
	ret = 1
	count := n //记录循环次数
	if n < 0 {
		count = -n
	}

	for i := 0; i < count; i++ {
		ret *= x
	}

	if n < 0 {
		return 1.0 / ret
	}

	return ret
}
