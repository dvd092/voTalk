package libs

import (
	"strings"
	"votalk/app/models"
)

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


//url最後の値取得
func SecondLastUrl(s string)  string {
	arr := UrltoSlice(s)
	lastar := arr[len(arr)-2]
	return lastar
}

//url最後の値短縮系
func LastUrlAb(s string)  string {
	arr := UrltoSlice(s)
	lastar := arr[len(arr)-1]
	if lastar == "viewer" {
		 lastar = "Vw"
	} else if lastar == "expert" {
		lastar = "Ex"
	}
	return lastar
}

//url最後の値int
func LastUrltoInt(s string)  int {
	arr := UrltoSlice(s)
	lastar := arr[len(arr)-1]
	var lastarI int
	if lastar == "viewer" {
		 lastarI = 2
	} else if lastar == "expert" {
		lastarI = 1
	}
	return lastarI
}

//テーブルのintからurl最後の値
func IntToLastUrl(n int) string {
	var lastarS string
	if n == 1 {
		 lastarS = "expert"
	} else if n == 2 {
		lastarS = "viewer"
	}
	return lastarS
}

// セッションからユーザータイプ取得
func GetUTypeFromSess(sess models.Session) (string) {
	return sess.UserType
}
