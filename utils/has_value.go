package utils

import "strings"

// HasValue 判断一个字符串是否有值
func HasValue(s string) bool {
	return strings.TrimSpace(s) != ""
}
