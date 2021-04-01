func solution(phone_number string) string {
    var phone_len = len(phone_number)
    var ret = ""
    last_four := string(phone_number[phone_len-4:phone_len])
    
    for i := 0; i < phone_len-4; i++{
        ret = ret + "*"
    }
    ret = ret + last_four
    return ret
}