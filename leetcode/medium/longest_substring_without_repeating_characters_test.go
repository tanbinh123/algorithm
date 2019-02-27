package medium

import "testing"

/**
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 */

 // 暴力法
func LengthOfLongestSubstring(s string) int {
	str := []rune(s)
	length := len(s)
	max := 0
	for i := 0; i < length; i++ {
		data := make(map[rune]int)
		jmax := 0
		for j := i; j < length; j++ {
			_, ok := data[str[j]]
			if ok {
				break
			}
			data[str[j]] = j
			jmax ++
		}
		if jmax > max {
			max = jmax
		}
	}
	return max
}




func TestLengthOfLongestSubstring(t *testing.T) {
	max := LengthOfLongestSubstring("abcabcbb")
	if max != 3 {
		t.Errorf("max error, result: %d", max)
	}
}
