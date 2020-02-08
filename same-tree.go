package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
type Pair struct {
	Left int
	Right int
}
func InOrder(node* TreeNode, serial *[]Pair, flag int){
    if node == nil {return}
	if node.Left != nil {InOrder(node.Left, serial, 1)}
	var elem = new(Pair)
	elem.Left = flag
	elem.Right = node.Val
	*serial = append(*serial,*elem)
	if node.Right!= nil {InOrder(node.Right, serial, 2)}
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var serial1 []Pair
	InOrder(p, &serial1, 0)
	var serial2 []Pair
	InOrder(q, &serial2, 0)
	
	
	if len(serial1) != len(serial2){return false;}
	for i:=0; i<len(serial1); i++ {
		var pair1 = serial1[i]
		var pair2 = serial2[i]
		if(pair1.Left != pair2.Left) {return false;}
		if(pair1.Right != pair2.Right){return false;}
	}

	return true;
}
func main() {
	var p = new (TreeNode)
	p.Val = 1
	p.Left = new (TreeNode)
	p.Left.Val = 1

	var q = new (TreeNode)
	q.Val = 1
	q.Right = new (TreeNode)
	q.Right.Val = 1
	println(isSameTree(p,q))
}