package main

/**
 * @Author: yirufeng
 * @Date: 2022/4/25 4:15 下午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/


//思路：使用两个变量记录横纵坐标，然后模拟就行了
func judgeCircle(moves string) bool {

	//记录横纵坐标
	x, y := 0, 0

	for _, ch := range moves {
		if ch == 'L' {
			x--
		} else if ch == 'R' {
			x++
		} else if ch == 'U' {
			y--
		} else {
			y++
		}
	}

	return x == 0 && y == 0
}