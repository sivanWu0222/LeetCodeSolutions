package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/21 4:05 下午
 * @Desc:
 **/

func hammingDistance(x int, y int) int {
	count := 0
	xorRet := x^y
	for xorRet != 0 {
		xorRet &= (xorRet-1)
		count++
	}
	return count
}