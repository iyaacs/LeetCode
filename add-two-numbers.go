package main

type ListNode struct {
	Val int
	Next *ListNode
}
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret = new(ListNode)
	var root = ret
	var len1 = 0
	var len2 = 0
	var allLen = 0
	var root1 = l1
	var root2 = l2
	for {
		len1++
		if l1.Next == nil {break}
		l1 = l1.Next		
	}
	for {
		len2++
		if l2.Next == nil {break}
		l2 = l2.Next
	}
	if len1 == len2 {allLen = len1}
	if len1>len2 {
		allLen=len1
		var diff = len1 - len2
		for i:=0; i<diff; i++ {
			l2.Next=new(ListNode)
			l2=l2.Next
			l2.Val=0
		}
	}
	if len1<len2 {
		allLen = len2
		var diff = len2 - len1
		for i:=0; i<diff; i++ {
			l1.Next=new(ListNode)
			l1=l1.Next
			l1.Val=0
		}
	}
	l1=root1
	l2=root2
	var plusVal = 0
	for i:=0; i<allLen; i++{
		var v1 = l1.Val
		var v2 = l2.Val
		var v3 = v1 + v2 + plusVal
		if v3 < 10 {plusVal = 0}
		if v3 >= 10 {
			plusVal = v3 / 10
			v3 = v3%10
		}
		ret.Val = v3
		if i==allLen-1 && plusVal > 0 {			
			ret.Next=new(ListNode)
			ret=ret.Next
			ret.Val = plusVal
		}
		if i < allLen-1 {
			l1=l1.Next
			l2=l2.Next
			ret.Next=new(ListNode)
			ret = ret.Next
		}
	}
	return root
}
func main() {
	var l1 = new(ListNode)
	var head1 = l1
	
	l1.Val=5

	var l2 = new(ListNode)
	var head2 = l2
	
	l2.Val=5
	var l3 = addTwoNumbers(head1, head2)
	for{		
		println(l3.Val)
		if l3.Next == nil {
			break
		}
		l3=l3.Next
	}

}