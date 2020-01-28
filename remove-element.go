package main

func removeElement(nums []int, val int) int {
	var left = 0
	for right:=0; right < len(nums); right++ {
		if nums[right] != val {
			var c = nums[right]
			nums[right] = nums[left]
			nums[left] = c
			left ++
		}
	} 
	return left
}
func main() {
	var nums = []int{0,1,2,2,3,0,4,2}
	println(removeElement(nums, 2))
	for _, val := range nums {
		print(val)
		print(" ")
	}
	println("")
	
}