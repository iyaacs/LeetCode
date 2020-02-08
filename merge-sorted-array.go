package main
import (
	"sort"
)
func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i:=m; i<m+n; i++{
		nums1[i]=nums2[i-m]
	}
	sort.Ints(nums1)
}
func main() {
	var nums1 = []int {1,2,3,0,0,0}
	var nums2 = []int {2,5,6}
	merge(nums1, 3, nums2, 3)
	println (len(nums1))
	for _, val := range(nums1){

		println (val)
	}
}