package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	//cmap中存储字符在字符串中的位置，位置不会超过int的范围
	cmap := make(map[uint8]int,128)
	//初始值不能为1，因为有可能输入的是空字符串
	ans := 0

	for i, j := 0, 0; j < n; j++ { //注意此处的i,j赋值方式
		if val,ok := cmap[s[j]];ok{
			i = val+1//从存在的字符后开始
		}
		cmap[s[j]] = j
		if j-i+1 > ans{
			ans = j-i+1
		}
	}
	return ans
}

func main(){
	s:= "abba"
	l := lengthOfLongestSubstring(s)
	fmt.Println(l)
}
