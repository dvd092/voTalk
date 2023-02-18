package libs

import "strings"

//url->スライス
func UrltoSlice(s string)  []string {
	arr := strings.Split(s, "/")
	return arr
}

//url最後の値取得
func LastUrl(s string)  string {
	arr := UrltoSlice(s)
	lastar := arr[len(arr)-1]
	return lastar
}