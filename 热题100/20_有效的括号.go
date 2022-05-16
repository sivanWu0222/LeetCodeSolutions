package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 9:28 下午
 * @Desc:
 **/

func isValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		//直接加入栈中
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}

		//直接从栈中弹出
		if len(stack) > 0 && ((s[i] == ')' && stack[len(stack)-1] == '(') ||
			(s[i] == ']' && stack[len(stack)-1] == '[') ||
			(s[i] == '}' && stack[len(stack)-1] == '{')) {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	//判断栈是否为空
	return len(stack) == 0
}
