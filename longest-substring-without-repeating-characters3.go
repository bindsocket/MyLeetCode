package leetcode

/*
Given a string s, find the length of the longest
substring
 without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

func lengthOfLongestSubstring(s string) int {
	var start int = 0
	var end int = 0
	var maxLength int = 0
	lastSeen := make(map[rune]int)
	for i, ThisChar := range s {
		if lastSeenIdx, ok := lastSeen[ThisChar]; ok {
			start = max(start, lastSeenIdx)
		}
		end = i
		lastSeen[ThisChar] = i + 1
		maxLength = max(maxLength, end-start+1)
	}
	return maxLength
}
