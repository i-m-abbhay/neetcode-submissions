func maxArea(heights []int) int {
    l, r := 0, len(heights)-1
    maxA := 0
    for l < r{
        maxA = max(maxA, (r-l)*min(heights[l], heights[r]))
        if heights[l]>heights[r]{
            r--
        } else {
            l++
        }
    }
    return maxA
}
