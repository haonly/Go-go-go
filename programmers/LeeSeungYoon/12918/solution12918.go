package Solution12918

import "unicode"

func Solution(s string) bool{

	// case 1: 길이가 4, 6인지 검사
	if len(s) != 4 && len(s) != 6{
		return false
	}

	// case 2: 모두 숫자인지 검사
	for _, runeVal := range s{
		if !unicode.IsDigit(runeVal){
			return false
		}
	}
	return true 
}
