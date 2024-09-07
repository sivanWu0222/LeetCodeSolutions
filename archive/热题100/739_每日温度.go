package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/24 12:22 下午
 * @Desc:
 **/

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return nil
	}

	stack := []int{}
	ret := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		//如果栈顶没有元素，直接加入当前元素
		if len(temperatures) == 0 {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			//负责结算
			ret[stack[len(stack)-1]] = i - stack[len(stack)-1]
			//弹出栈顶
			stack = stack[:len(stack)-1]
		}

		//加入当前元素
		stack = append(stack, i)
	}
	return ret
}
