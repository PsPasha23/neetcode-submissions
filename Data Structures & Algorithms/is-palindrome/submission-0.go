func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r{
		fmt.Println("left: %v, right: %v", string(s[l]), string(s[r]))
		if !isAlphaNum(s[l]){
			l++
		} else if !isAlphaNum(s[r]){
			r--
		} else if strings.ToLower(string(s[l])) == strings.ToLower(string(s[r])){
			l++
			r--
		}   else{
			return false
		}
	}
	return true
}
func isAlphaNum (c byte) bool{
	return (c >= 'a' && c <= 'z') || ( c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
