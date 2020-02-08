package main

type ListNode struct {
	Val int
	Next *ListNode
 }
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {return nil;}
	if head.Next == nil {return head;}
	var root = head
	for{
		if head.Next == nil {break;}
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
			continue
		}		
		head = head.Next
	}
	return root;
}
func main() {
	var root = new (ListNode)
	var elem = root
	var arr = []int {1,1,2,3,3}
	for _, val := range(arr){
		elem.Val = val
		elem.Next = new (ListNode)
		elem = elem.Next
	}
	var ret = deleteDuplicates(root)
	for{
		if ret.Next == nil {break}
		println(ret.Val)
		ret = ret.Next
	}
}