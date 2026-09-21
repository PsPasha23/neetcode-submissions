func twoSum(nums []int, target int) []int {
    antiMap := make(map[int]int, len(nums))
	for i:= 0; i<len(nums); i++{
		if val, ok := antiMap[nums[i]]; ok{
			return []int{val, i}
		}
		antiMap[target-nums[i]] = i
	}
	return []int{}
}
