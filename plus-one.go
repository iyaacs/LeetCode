
package main

func plusOne(digits []int) []int {
	var ret []int
	digits[len(digits)-1]++
	var up = 0
	for i := len(digits)-1; i>=0; i-- {
		var val = digits[i] + up
		
		if val >= 10 {
			up=1
			val %= 10
		} else {up = 0}
		ret = append(ret, val)
	}
	if up > 0 {ret = append(ret, up)}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return ret
}

func main() {
	var digits = []int {4,3,2,1,9}
	var ret = plusOne(digits)
	
	for _, val := range(ret) {
		println (val)
	}
}