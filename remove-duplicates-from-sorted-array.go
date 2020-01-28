package main

func removeDuplicates(nums []int) int {	
    if len(nums) <= 1 {
		return len(nums)
	}
	var left = 0

	for right := 1; right < len(nums); right++ {
		if nums[left] >= nums[right] {continue}
		left++
		nums[left]=nums[right]
	}
	return left+1	
}
func main() {
	var nums = []int{0,1}
	println(removeDuplicates(nums))
	for _, val := range nums {
		print(val)
		print(" ")
	}
	println("")
	
}