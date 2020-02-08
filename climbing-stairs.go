package main


func climbStairs(n int) int {
	if n == 0 {return 0;}
	var v1 = 0
	var v2 = 1
	var v3 = 1
	for i:=0; i<n; i++{
		v3=v1+v2
		v1=v2
		v2=v3
	}
	return v3;
}

func main() {
	var num = 1
	println(climbStairs(num))
}