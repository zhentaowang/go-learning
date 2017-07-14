package utils

import "strings"

// switch format from hump to snake
func FormatSwitch(key string) (string) {

	for _, char := range key {
		if char >= 'A' && char <= 'Z' {
			key = strings.Replace(key, string(char), "_" + strings.ToLower(string(char)), -1)
		}
	}
	return key

}
