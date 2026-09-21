func groupAnagrams(strs []string) [][]string {
    stringMapping := make(map[[26]int][]string, len(strs))
    for _, str := range strs{
        var patternMapping [26]int
        for i := 0; i < len(str); i++{
            patternMapping[str[i]-'a'] += 1
        }
        if val, ok :=  stringMapping[patternMapping] ; ok{
            val = append(val, str)
            stringMapping[patternMapping] = val
        } else{
            stringMapping[patternMapping] = []string{str}
        }
    }
    response := make([][]string, 0)
    for _, val := range stringMapping{
        oneValStrings := []string{}
        oneValStrings = append(oneValStrings, val...)
        response  = append(response, oneValStrings)
    }
    return response
}
