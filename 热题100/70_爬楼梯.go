package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 4:25 下午
 * @Desc:
 **/


//第一种：递归解法
//func climbStairs(n int) int {
//	if n <= 0 {
//		return 0
//	}
//
//	if n <= 2 {
//		return n
//	}
//
//	return climbStairs(n-1) + climbStairs(n-2)
//}
//

//第二种：迭代解法
func climbStairs(n int) int {
	if n <= 0 {
		return 0
	}

	if n <= 2 {
		return n
	}

	one, two := 1, 2

	for i := 0; i < n-2; i++ {
		one, two = two, one+two
	}

	return two
}