package leetcode

//problem link [https://leetcode.com/problems/contains-duplicate/description/]

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}
