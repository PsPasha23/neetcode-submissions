type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	lenStr := ""
	appendStr := ""
	for _, str := range strs{
		leng := strconv.Itoa(len(str))
		lenStr = lenStr + leng + ","
		appendStr = appendStr + str
	}

	return lenStr+"#" +appendStr
}

func (s *Solution) Decode(encoded string) []string {

	countArr := make([]int, 0)
	actualStr := ""
	numStr := ""
	for i, s := range encoded{
		if s == '#'{
			actualStr = encoded[i+1:]
			break
		}
		if s == ','{
			a, _ := strconv.Atoi(numStr)
			countArr = append(countArr, a)
			numStr = ""
			continue
		}
		numStr += string(s)
	}
	res := make([]string, 0, len(countArr))
	prev := 0
	for _, i := range countArr{

		res = append(res, actualStr[prev:prev+i])
		prev += i
	}
	return res
}
