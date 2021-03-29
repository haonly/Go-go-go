func solution(n int) int64 {
    var mod int = 1234567 
    var dp = make([]int, 0, 3)

    for i := 0; i < n; i++{
        if len(dp) < 3{
            dp = append(dp, i+1)
        }else{
            dp = append(dp, dp[i-1] % mod + dp[i-2] % mod)
        }
    }
    return int64(dp[len(dp)-1] % mod)
}



// func solution(n int) int64 {
    
//     var dp [2001]int
//     dp[0]=1
//     dp[1]=1
//     for i:=2; i<=n; i++ {
//         dp[i]= (dp[i-2] + dp[i-1])%1234567
//     }

//     return int64(dp[n])
// }