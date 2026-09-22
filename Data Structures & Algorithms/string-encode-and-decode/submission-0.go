type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	res := ""
	for _, str := range strs{
		res = res + "3_" + str
	}
	return res
}

func (s *Solution) Decode(encoded string) []string {
	strs := strings.Split(encoded, "3_")
	return strs[1:]
}
