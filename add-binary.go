package main

func reverseString(s string) string {
	var ret = ""
	for i := len(s)-1; i>=0; i-- {
		ret += string(s[i])
	}
	return ret
}

func paddingString(s string, x int) string {
	var ret = ""
	for i := len(s); i<x; i++ {
		ret+="0"
	}
	for i:=0; i<len(s); i++ {
		ret+=string(s[i])
	}
	return ret
}

func addBinary(a string, b string) string {
	var l1 = len(a)
	var l2 = len(b)

	var maxLen = l1
	if l2 > maxLen {maxLen = l2}

	a = paddingString(a, maxLen)
	b = paddingString(b, maxLen)
	var revA = reverseString(a)
	var revB = reverseString(b)

	var ret = ""
	var up = 0
	for i:=0; i<maxLen; i++ {
		var v1 = int(revA[i]) - 48
		var v2 = int(revB[i]) - 48
		var ssum = v1 + v2 + up
		up = ssum / 2
		ret += string((ssum%2)+48)
		
	}
	if up == 1 {
		ret += "1"
	}	

	return reverseString(ret)
}
func main() {
	var a = "11"
	var b = "1"
	println(addBinary(a,b))
}