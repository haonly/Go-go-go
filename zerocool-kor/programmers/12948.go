/*
핸드폰 번호 가리기
Description
프로그래머스 모바일은 개인정보 보호를 위해 고지서를 보낼 때 고객들의 전화번호의 일부를 가립니다.
전화번호가 문자열 phone_number로 주어졌을 때, 전화번호의 뒷 4자리를 제외한 나머지 숫자를 전부 *으로 가린 문자열을 리턴하는 함수, solution을 완성해주세요.

제한 조건
s는 길이 4 이상, 20이하인 문자열입니다.
입출력 예
phone_number	return
"01033334444"	"*******4444"
"027778888"	"*****8888"
*/

package main

import (
	"errors"
	"log"
)

func solution(s string) (string, error) {
	log.Println("solution s:", s)

	if len(s) >= 4 && len(s) <= 20 {
		rs := []rune(s)
		for i := 0; i < len(rs)-4; i++ {
			rs[i] = '*'
		}
		return string(rs), nil
	} else {
		return "", errors.New("sting length error (4 <= str <= 20)")
	}
}

func main() {
	s1, _ := solution("01033334444")
	log.Println("result -> ", s1)
	s2, _ := solution("027778888")
	log.Println("result -> ", s2)
	s3, _ := solution("5190")
	log.Println("result -> ", s3)
}
