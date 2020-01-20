package main
import (
	"math"
)
func GetNumberPlace(x int, n int) int {
	var len = int(math.Log10(float64(x)))+1
	var idx1 = len - 1 - n
	var pow1 = int(math.Pow10(idx1))
	var x1 = x / pow1
	x1 *= pow1

	var idx2 = len - n
	var pow2 = int(math.Pow10(idx2))
	var x2 = x / pow2
	x2 *= pow2

	var v1 = x - x1
	var v2 = x - x2
	var v3 = v2 - v1
	if v3 == 0 {
		return 0
	}
	var lenv = int(math.Log10(float64(v3)))
	var powv = int(math.Pow10(lenv))

	return v3 / powv
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var lenx = int(math.Log10(float64(x)))+1
	var left = 0
	var right = lenx - 1
	for left < right {
		var v1 = GetNumberPlace(x, left)
		var v2 = GetNumberPlace(x, right)
		if v1 != v2{
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	var x = 12321
	//println(GetNumberPlace(x,5))
	println(isPalindrome(x))
}