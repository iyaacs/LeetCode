package main
func CtoI(s string) int {
	if s=="I"{return 1}
	if s=="V"{return 5}
	if s=="X"{return 10}
	if s=="L"{return 50}
	if s=="C"{return 100}
	if s=="D"{return 500}
	if s=="M"{return 1000}
	return 0
}
func romanToInt(s string) int {
	var ret = 0
	s = s + "0"

	var flag = false;
	for idx :=0; idx< len(s); idx++ {
		if flag==true {flag=false; continue;}
		var i1 = CtoI(string(s[idx]))
		if idx == len(s)-1 {break;}
		
		var i2 = CtoI(string(s[idx+1]))
		if i1 >= i2 {
			ret += i1;
		} else {
			ret += i2; ret-= i1; flag=true;
		}
	}
	return ret
}
func main() {
	var x = "IX"
	//println(GetNumberPlace(x,5))
	println(romanToInt(x))
}