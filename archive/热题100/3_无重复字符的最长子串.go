package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 4:21 下午
 * @Desc:
 **/

func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	ret := 0
	hmap := make(map[byte]bool)

	//Tips: 这里是 r < len(s) 不是 l <= r
	for r < len(s) {
		//如果当前要加入的字符存在于窗口中，不断从窗口左侧移出，直到没有当前要加入的字符
		if _, ok := hmap[s[r]]; ok {
			delete(hmap, s[l])
			l++
			continue
		}

		//此时将当前字符加入到我们的窗口中
		hmap[s[r]] = true
		//r后移
		r++
		//更新我们的结果
		if ret < r-l {
			ret = r - l
		}
	}
	return ret
}
