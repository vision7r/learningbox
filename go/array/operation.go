package main

import "fmt"

// array elements store in the memory one after one
// array is value type
// time complexity - lookup and append by index O(1)
// Insert at[2], search , delete[3] O(n).
func main() {
	//Declaration by var keyword
	var nums1 [5]int
	nums1 = [5]int{1, 2, 3, 4, 5}
	nums2 := [6]int{1, 2, 3, 4}
	//nums2=[1,2,3,4,0,0] here index 4,5 is 0 cause, by default an array is zero-valued .
	fmt.Println(nums1, nums2)
	//lookup at specific index
	fmt.Println(nums1[3])

	//Insert at index[2] in nums2 example
	// for this just move to the number right side
	length := 4
	for i := length; i > 2; i-- {
		nums2[i] = nums2[i-1]
	}
	nums2[2] = 10
	length++
	fmt.Println(nums2)
	fmt.Println("Now length is", length)

	//Delete at  index[2] in nums2 example
	// for this we have to do opposite move to left side
	for i := 2; i < length; i++ {
		nums2[i] = nums2[i+1]
	}
	length--
	fmt.Println(nums2)
	fmt.Println("Now length is", length)

}
