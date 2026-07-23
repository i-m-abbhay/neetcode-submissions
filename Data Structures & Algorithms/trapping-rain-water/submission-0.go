func trap(height []int) int {
    sz := len(height)
    maxFromLeft := make([]int, sz)
    maxFromRight := make([]int, sz)
    totalWater := 0
    maxFromLeft[0]= height[0]
    maxFromRight[sz-1] = height[sz-1]
    maxFromL := height[0]
    maxFromR := height[sz-1]
    for i:= 1;i<sz;i++{
        maxFromLeft[i] = max(maxFromL, height[i])
        maxFromL = maxFromLeft[i]
    }
    for i:= sz-2;i>=0;i--{
        maxFromRight[i] = max(maxFromR, height[i])
        maxFromR = maxFromRight[i]
    }
    for i:= 0; i<sz;i++{
        totalWater += min(maxFromLeft[i], maxFromRight[i])-height[i]
    }
    return totalWater
}
