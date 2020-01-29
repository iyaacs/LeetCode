package main
import (
	"strconv"
)
func GetNextString(str *string) string {
	if len(*str)<1 {
		return "1"
	}
	(*str)+="_"

	var ret = ""
	var char1 = (*str)[0]
	var cnt = 0
	for i:=0; i<len(*str); i++ {
		var char2 = (*str)[i]
		if string(char1) == string(char2) {
			cnt++
			continue
		} else {
			ret+=strconv.Itoa(cnt)
			ret+=string(char1)
			char1=char2
			cnt=1
		}
	}

	(*str)=(*str)[0:len(*str)-1]
	return ret
}
func Generate(arr *[]string) {
	if (*arr) != nil {
		if len(*arr) > 0 {
			return
		}
	}
	var newArr []string
	var s = ""
	for i:=0; i<30; i++{
		s = GetNextString(&s)
		newArr = append(newArr, s)
	}
	*arr = newArr
	return 
}
var cache[] string
func countAndSay(n int) string {
	Generate(&cache)
    return cache[n-1]
}
func main() {
	for i:=1; i<=30; i++ {
		println(countAndSay(i))
	}
}