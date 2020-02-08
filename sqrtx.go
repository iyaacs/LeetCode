package main
import (
	"math"
)

func mySqrt(x int) int {
    return int(math.Sqrt(float64(x)))
}

func main() {
	var num = 8
	println(mySqrt(num))
}