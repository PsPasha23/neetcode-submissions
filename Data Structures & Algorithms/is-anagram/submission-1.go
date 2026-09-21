func isAnagram(s string, t string) bool {
	compArr := make([]int, 26)
	ascciVala := int('a')
	for _, c := range s{
		asciiVal := int(c)
		compArr[asciiVal-ascciVala] += 1
	}
	for _, c := range t{
		asciiVal := int(c)
		compArr[asciiVal-ascciVala] -= 1
	}
	for _, a := range compArr{
		if a != 0{
			return false
		}
	}
	return true
}
