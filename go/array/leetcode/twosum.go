package leetcode

import "fmt"

//problem link {https://leetcode.com/problems/two-sum/description/}

func twosum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		compliment := target - num

		if idx, yes := m[compliment]; yes {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	s := []int{1, 3, 5, 6, 4}
	t := 9
	fmt.Println("Answer is :", twosum(s, t))
}
