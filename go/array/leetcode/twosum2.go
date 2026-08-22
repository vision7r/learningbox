package leetcode

//problem link {https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/}

func twosumd(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		compliment := target - num

		if idx, yes := m[compliment]; yes {
			return []int{idx + 1, i + 1}
		}
		m[num] = i
	}
	return nil
}

/*func main(){
	s := []int{1,3,5,6,4}
	t := 9
	fmt.Println("Answer is :",twosumd(s,t))
}*/
