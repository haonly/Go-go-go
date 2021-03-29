import (
    "strconv"
)

func solution(n int) int {
    answer := 0
    
    var str_n string = strconv.Itoa(n)
    for i := 0; i < len(str_n); i++ {
        val, _ := strconv.Atoi(string(str_n[i]))
        answer += val
    }
//     for {
//         answer += n % 10
//         n /= 10

//         if n == 0 {
//             break
//         }
//     }

    return answer
}