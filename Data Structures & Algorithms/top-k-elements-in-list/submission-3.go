func topKFrequent(nums []int, k int) []int {
	
	count := make(map[int]int, 0)
	for _, n := range nums{
		count[n] += 1
	}
	frequencyMap := make(map[int][]int)
	for key,val := range count {
		frequencyMap[val] = append(frequencyMap[val], key)
	}
	res := make([]int, 0,k)
	for i := len(nums) ; len(res) < k && i>0; i--{
		res = append(res, frequencyMap[i]...)
	}
	return res
}
