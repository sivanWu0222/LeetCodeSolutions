package main

/**
 * @Author: yirufeng
 * @Date: 2022/4/28 6:53 下午
 * @Email: yirufeng@foxmail.com
 * @GitHub: https://www.github.com/sivanWu0222
 * @Desc:
 **/
func uniquePaths(m int, n int) int {

	//基本的动态规划
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	//初始化第一行
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	//初始化第一列
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

//方法二：进行状态压缩之后的
func uniquePaths(m int, n int) int {

	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			ret[j] = ret[j-1] + ret[j]
		}
	}

	return ret[n-1]
}
