package leetcode

//problem link [https://leetcode.com/problems/product-of-array-except-self/description/]

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// answer[i] will first hold the product of all elements before i
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	// now multiply by the product of all elements after i,
	// tracked with a running suffix product variable
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}

	return answer
}
