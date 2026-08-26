package leetcode

//problem Link [https://leetcode.com/problems/maximum-subarray/description/]

//for this follow Kadane's algorithm

func maxSubArray(nums []int) int {
	cu_sum := nums[0]
	maxt := nums[0]

	for i := 1; i < len(nums); i++ {
		cu_sum = max(nums[i], cu_sum+nums[i])
		maxt = max(maxt, cu_sum)

	}
	return maxt
}
