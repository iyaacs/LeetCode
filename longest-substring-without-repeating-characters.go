package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {return 0}
	cache := make(map[string]int)
	var left = 0
	var right = 0
	var max = 0
	var now = 0
	// 0 : right go, 1 : left go
	var flag = 0
	for {
		if flag == 0 {
			var c = string(s[right])
			if cache[c] == 0 {
				cache[c]++
				now++
				if now>max {max=now}
				right++
				if right >= len(s) {break}
			} else {
				flag = 1
				continue
			}
		}
		if flag == 1 {
			var c = string(s[left])
			left++
			now--
			cache[c]--
			if cache[c] == 0 {
				flag=0
				continue
			}
		}
	}
	return max
}
func main(){

	println(lengthOfLongestSubstring("abcabcbb"))
}