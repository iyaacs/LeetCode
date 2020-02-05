
package main

func maxSubArray(nums []int) int {
	if nums == nil {return 0;}
	if len(nums) == 0 {return 0;}
	if len(nums) == 1 {return nums[0];}

	var maxSum = nums[0];
	var nowSum = nums[0];
	for i:=1; i<len(nums); i++ {
		var v = nums[i];
		if nowSum < 0 {nowSum = 0;}
		nowSum += v;
		if nowSum > maxSum {
			maxSum = nowSum;
		}
	}
    return maxSum;
}

func main() {
	var nums = []int {-2,1,-3,4,-1,2,1,-5,4}
	println(maxSubArray(nums))
}