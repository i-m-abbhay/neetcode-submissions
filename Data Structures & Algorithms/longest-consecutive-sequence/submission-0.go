func longestConsecutive(nums []int) int {
    numSet := make(map[int]bool)
    for _, num:= range nums{
        numSet[num]=true
    }
    longest:=0

    for num:=range numSet{
        currLen := 0
        if !numSet[num-1]{
            currLen++
            for numSet[num+1]{
                currLen++
                num++
            }
            longest = max(currLen, longest)
        }
    }
    return longest
}
