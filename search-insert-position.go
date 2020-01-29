package main
func LowerBound(nums[] int, target int, leftIndex int, rightIndex int) int {
	if leftIndex > rightIndex {
		return leftIndex
	}
	midIndex:=(leftIndex+rightIndex)/2
	if target == nums[midIndex] {
		return midIndex
	} else if target < nums[midIndex] {
		return LowerBound(nums, target, leftIndex, midIndex-1)
	}
	return LowerBound(nums, target, midIndex+1, rightIndex)
}
func searchInsert(nums []int, target int) int {
    return LowerBound(nums, target, 0, len(nums)-1)
}
func main() {
	var nums = []int {1,3,5,6}
	var target = 0
	println(searchInsert(nums,target))
	
}