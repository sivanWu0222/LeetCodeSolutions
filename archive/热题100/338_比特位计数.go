package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/20 11:12 上午
 * @Desc:
 **/

func countBits(n int) []int {
	ret := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		ret[i] = ret[i&(i-1)] + 1
	}
	return ret
}
