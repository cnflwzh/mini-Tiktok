package utils

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// HashPassword 生成加盐的哈希值
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash 检查密码是否匹配
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CheckUsername 检查用户名是否合法
func CheckUsername(username string) bool {
	// 匹配3-20个字符，包括小写字母、大写字母、数字或下划线
	reg, err := regexp.Compile("^[a-zA-Z0-9_]{3,20}$")
	if err != nil {
		hlog.Error("Error compiling regex: %v", err)
		return false
	}
	return reg.MatchString(username)
}

// CheckPassword 检查密码是否合法
func CheckPassword(password string) bool {
	uppercase := regexp.MustCompile(`[A-Z]+`)
	lowercase := regexp.MustCompile(`[a-z]+`)
	number := regexp.MustCompile(`[0-9]+`)

	return len(password) >= 6 && len(password) <= 20 &&
		uppercase.MatchString(password) &&
		lowercase.MatchString(password) &&
		number.MatchString(password)
}

// TrimSpace 去除字符串两端的空格
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// TimestampToFormatTime 将时间戳转换为格式化的时间
func TimestampToFormatTime(timestamp string) string {
	// 将字符串转换为整数
	ms, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		fmt.Printf("Error parsing timestamp to int: %v\n", err)
		return ""
	}

	// 将毫秒为单位的Unix时间戳转换为time.Time类型
	t := time.Unix(0, ms*int64(time.Millisecond))

	// 获取时间的格式化字符串
	formattedTime := t.Format("2006-01-02 15:04:05")

	return formattedTime
}
