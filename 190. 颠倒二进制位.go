/**
 @Author       cvenwu
 @Datetime     2022/9/26 22:15
 @Description
**/
package main

func reverseBits(num uint32) uint32 {
	var ret uint32
	ret = 0
	for i := 0; i < 32; i++ {
		// 首先将结果向左移动1位，然后加上原始数字第i位的值(判断是否为1，如果为1就加上)
		ret = ret<<1 + (num>>i)&1
	}
	return ret
}
