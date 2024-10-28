package utils

import "strings"

func Split2(s string) (string, string) {
	arr := strings.SplitN(s, " ", 2)
	return arr[0], arr[1]
}
