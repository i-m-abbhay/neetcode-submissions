type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    var sb strings.Builder
    for _, str := range strs {
       sb.WriteString(strconv.Itoa(len(str)))
       sb.WriteString("#")
       sb.WriteString(str)
    }
    return sb.String()

}

func (s *Solution) Decode(str string) []string {
    res := []string{}
    i:= 0
    for i<len(str){
        j:= strings.Index(str[i:], "#")
        hashPos := i+j
        lengthStr := str[i:hashPos]
        length, _ := strconv.Atoi(lengthStr)
        content := str[hashPos+1:hashPos+1+length]
        i=hashPos + 1 + length
        res = append(res, content)
    }
    return res
}
