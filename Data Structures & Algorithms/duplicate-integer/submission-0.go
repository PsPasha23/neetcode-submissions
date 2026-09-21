func hasDuplicate(nums []int) bool {
    frequencyMap := make(map[int]bool, len(nums))
    for _, n := range nums{
        if _, ok := frequencyMap[n]; ok{
            return true
        }
        frequencyMap[n] = true
    }
    return false
}
