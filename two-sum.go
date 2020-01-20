

package main

func twoSum(nums []int, target int) []int {
	for index1, value1:= range nums {
		for index2, value2:= range nums{
			if index2 <= index1 {
				continue
			}
			if value1+value2 == target{
				var ret = []int{index1, index2}
				return ret
			}
		}
	}
	var ret = []int{0, 0}
	return ret
}
func main() {
	var arr = []int{2, 7, 11, 15}
	var ret = twoSum(arr, 9)
	println(ret[0])
	println(ret[1])
}