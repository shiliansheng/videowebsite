package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

// ### 生成唯一ID

// 生成唯一ID
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return Md5(base64.URLEncoding.EncodeToString(b))
}

func Md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str) )
	return fmt.Sprintf("%x", hash.Sum(nil))
}




func TimePareStd(value string) (time.Time, error) {
	ret, err := time.Parse("2006-01-02 15:03:04", value)
	return ret, err
}

const (
	TimeFormart = "2006-01-02 15:04:05"
)

func GetNowTimeString() string {
	return time.Now().Format(TimeFormart)
}

func Atoi(s string) int {
	res, _ := strconv.Atoi(s)
	return res
}