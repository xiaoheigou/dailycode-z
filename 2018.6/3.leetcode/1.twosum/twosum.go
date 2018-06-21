package twosum

func twoSum(nums []int, target int) []int {
	//m 负责保存map[序号]
	m := make(map[int]int)

	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
