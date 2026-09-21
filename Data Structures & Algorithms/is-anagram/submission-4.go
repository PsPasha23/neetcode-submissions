func isAnagram(s string, t string) bool {
	compArr := make([]int, 26)
	if len(s)!= len(t){
		return false
	}

	for i := 0; i< len(s); i++{
		compArr[s[i]-'a'] += 1
		compArr[t[i]-'a'] -= 1
	}
	
	for _, a := range compArr{
		if a != 0{
			return false
		}
	}
	return true
}
