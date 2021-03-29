package ValidParentheses_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ValidParentheses "leetcode/20_Valid_Parentheses"
)

func TestCase1(t *testing.T) {
	var s string = "()"
	answer := true
	testResult := ValidParentheses.IsValid(s)
	assert.Equal(t, testResult, answer, "Wrong Answer")
}

func TestCase2(t *testing.T) {
	var s string = "()[]{}"
	answer := true
	testResult := ValidParentheses.IsValid(s)
	assert.Equal(t, testResult, answer, "Wrong Answer")
}

func TestCase3(t *testing.T) {
	var s string = "([)]"
	answer := false
	testResult := ValidParentheses.IsValid(s)
	assert.Equal(t, testResult, answer, "Wrong Answer")
}

func TestCase4(t *testing.T) {
	var s string = "{[]}"
	answer := true
	testResult := ValidParentheses.IsValid(s)
	assert.Equal(t, testResult, answer, "Wrong Answer")
}
