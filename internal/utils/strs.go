package utils

import "strings"

// NotEmpty 判断字符串是否为空，长度为零
func NotEmpty(str string) bool {
	return len(str) > 0
}

// NotBlank 判断是否为空字符串
func NotBlank(str string) bool {
	return len(strings.TrimSpace(str)) > 0
}
