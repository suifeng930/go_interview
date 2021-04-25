package main

import (
	"strings"
)

// 通过 steings.count判断
// strings.Count  返回字符串s 中有杰哥不重复的sep 子串
// 如果subStr 是一个空字符串,则 count 返回(1+s)中的unicode代码点数
func isUniqueString	(s string) bool  {

	if strings.Count(s,"")>3000 {
		return false
	}

	for _, v := range s {
		if v>127 {
			return false
		}
		if strings.Count(s,string(v))>1 {
			return false
		}
	}
	return true
}

// strings.Index 返回s 中substr的第一个实例的索引，如果不存在返回-1
func isUniqueString2(s string) bool {

	if strings.Count(s,"")>3000 {
		return false
	}
	for k, v := range s {
		if v>127 {
			return false
		}
		if strings.Index(s,string(v)) !=k {
			return false
		}
	}
	return false
}