package ValidParentheses

type Stack []byte

func isContained(b byte, list []byte) bool {
	for _, v := range list {
		if b == v {
			return true
		}
	}
	return false
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(b byte) {
	*s = append(*s, b)
}

func (s *Stack) pop() (byte, bool) {
	if s.isEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) head() (byte, bool) {
	if s.isEmpty() {
		return 0, false
	} else {
		return (*s)[len(*s)-1], true
	}
}

func IsValid(s string) bool {

	var sLength int = len(s)

	// 홀수면 false 리턴
	if sLength%2 == 1 {
		return false
	}

	frontParentheses := []byte{'(', '{', '['}
	dict := map[byte]byte{
		'(': ')',
		')': '(',
		'{': '}',
		'}': '{',
		'[': ']',
		']': '[',
	}
	var stack Stack

	for i := 0; i < sLength; i++ {
		if isContained(s[i], frontParentheses) {
			// 문자가 front 괄호인 경우
			stack.push(s[i])
		} else {
			// 문자가 end 괄호인 경우
			head, flag := stack.head()
			// head가 존재하는 경우
			if flag {
				if dict[head] != s[i] { // head와 짝꿍 괄호가 아닌 경우
					return false
				} else {
					stack.pop()
				}
			} else { // end 괄호가 먼저 나온 경우
				return false
			}
		}
	}
	return stack.isEmpty() // 비어있으면 true
}
