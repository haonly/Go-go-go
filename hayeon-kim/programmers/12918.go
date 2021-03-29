func solution(s string) bool {
    if len(s) == 4 || len(s) == 6{
        for i := 0; i < len(s); i++ {
            if 48 <= s[i] && s[i] <= 57 {
                continue
            } else{
                return false
            }
        }
    }else{
        return false
    }
    return true
}