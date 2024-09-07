package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/20 12:14 下午
 * @Desc:
 **/

func rotate(matrix [][]int) {

	//首先获取矩阵的维度
	n := len(matrix)

	//大的循环次数
	for i := 0; i < n>>1; i++ {
		//小的循环次数
		for j := i; j < n-1-i; j++ {
			//左上，右上，右下，左下 = 右上，右下，左下, 右上
			matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
		}
	}
}