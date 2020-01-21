package main

type Node struct {
	c string
	IsTerminal bool
	nextPtr *Node
}

type Tri struct {
	root *Node
}

func GetTri() *Tri {
	var ret = new(Tri)
	ret.root = new(Node)
	ret.root.c = ""
	ret.root.IsTerminal=false
	ret.root.nextPtr=nil
	return ret
}

func (tri *Tri) Insert(s string) {
	var node = tri.root
	for _,c := range s{
		if node.IsTerminal==true {
			return
		}
		if node.nextPtr == nil {
			node.nextPtr = new(Node)
			node = node.nextPtr
			node.c = string(c)
			node.IsTerminal = false
			node.nextPtr = nil
			continue
		}
		if node.nextPtr.c == string(c) {
			node = node.nextPtr
			continue
		} else {
			node.IsTerminal = true
			return
		}
	}			
	node.IsTerminal = true
	return
}
func (tri *Tri) GetResult () string {
	var ret = ""
	var node = tri.root
	for {
		if len(node.c)>0 {
			ret += node.c
		}
		if node.IsTerminal == true {
			break
		}
		node = node.nextPtr
	}
	return ret
}


func longestCommonPrefix(strs []string) string {
	if (len(strs)==0) {return ""}
	var tri = GetTri()
	for i:=0; i<len(strs); i++ {
		tri.Insert(strs[i])
	}
	return tri.GetResult()
}

func main() {
	var strs = []string{}
	println(longestCommonPrefix(strs))
	
}