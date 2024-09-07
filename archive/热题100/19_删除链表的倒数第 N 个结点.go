package 热题100

/**
 * @Author: yirufeng
 * @Date: 2021/10/19 4:40 下午
 * @Desc:
 **/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//因为可能删除头节点，所以需要建立一个虚拟节点
	dummyNode := &ListNode{}
	dummyNode.Next = head

	//快指针先走n+1步
	fast, slow := dummyNode, dummyNode
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	//之后快慢一起走，当快走到nil的时候，慢指向倒数第n个节点的前一个节点
	for fast != nil {
		fast, slow = fast.Next, slow.Next
	}

	//此时修改慢指针的Next指向
	slow.Next = slow.Next.Next
	return dummyNode.Next
}
