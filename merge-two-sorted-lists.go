

package main

type ListNode struct {
	Val int
	Next *ListNode
}
 
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {return l2}
	if l2 == nil {return l1}

	var ret = new (ListNode)
	var root = ret
	for{
		var flag = 0
		if l1 == nil && l2 == nil {
			break
		}
		var tmp = new (ListNode)
		if l1 == nil {
			tmp = l2
			flag=2
		} else if l2 == nil {
			tmp = l1
			flag=1
		} else {
			if l1.Val < l2.Val {
				tmp = l1
				flag=1
			} else {
				tmp = l2
				flag=2
			}
		}
		ret.Next = tmp
		ret = ret.Next
		tmp = tmp.Next

		if flag == 1 {l1 = tmp}
		if flag == 2 {l2 = tmp}
	}
	ret = root.Next
	return ret
}
func main(){
	var l1 = new (ListNode)
	var r1 = l1
	l1.Val = 1
	l1.Next = new (ListNode)
	l1 = l1.Next

	l1.Val = 2
	l1.Next = new (ListNode)
	l1 = l1.Next

	l1.Val = 4

	var l2 = new (ListNode)
	var r2 = l2
	l2.Val = 1
	l2.Next = new (ListNode)
	l2 = l2.Next

	l2.Val = 3
	l2.Next = new (ListNode)
	l2 = l2.Next
	
	l2.Val = 4
	
	l1 = r1
	l2 = r2
	
	var node = mergeTwoLists(l1, l2)

	for{
		if node == nil {break}
		println(node.Val)
		node = node.Next
	}
}