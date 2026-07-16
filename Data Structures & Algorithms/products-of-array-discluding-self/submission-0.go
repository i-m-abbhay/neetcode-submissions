func productExceptSelf(nums []int) []int {
    //1 2 4 6
   n:= len(nums)

   res:= make([]int, n)
   res[0] = 1
    //1 0 0 0
    for i:=1;i<n;i++{
        res[i]=res[i-1]*nums[i-1]
        //1 1 2 8
    }
    sufSoFar := 1
    for i:=n-1;i>=0;i--{
        res[i] = sufSoFar*res[i] // 1 1 2 8
        sufSoFar = nums[i]*sufSoFar 
    }

   return res

}
