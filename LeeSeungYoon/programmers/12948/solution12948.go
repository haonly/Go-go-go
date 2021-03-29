//https://programmers.co.kr/learn/courses/30/lessons/12948

package Solution12948

func Solution(phone_number string) string{

	var length int = len(phone_number)
	var answer string = ""

	for i:=0; i<length-4; i++{
		answer += "*"
	}
	answer += phone_number[length-4:]

	return answer
}
