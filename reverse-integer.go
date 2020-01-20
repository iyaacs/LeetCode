
package main
import (
	"math"
)
func reverse(x int) int {


	var sign = 1
	if x < 0 {
		sign = -1
	}
	var absx = x * sign
	var pow10 = int(math.Pow(10, float64(int(math.Log10(float64(absx))))))
	var ret = 0

	for absx > 0 {
		var absx2 = absx / 10;
		absx2 *= 10
		var n = absx - absx2
		var k = (n*pow10)
		ret += k
		absx = absx2 / 10
		pow10 /= 10
	} 
	
	var realret = ret * sign
	if realret < - int(math.Pow(2, 31)){
		return 0;
	}
	if realret > int(math.Pow(2, 31))-1{
		return 0;
	}
	return realret;
}

func main() {
	var x = 534236469
	println(reverse(x))
}