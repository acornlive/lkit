package strkit

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	formatTime     = "2006-01-02 15:04:05"
	whiteSpace     = ' '
	seedStr        = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	regDigit       = "^[0-9]*$"
	regTime        = "^(?:[01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d$"
	regDateTime    = "^(\\d{1,4}-(?:1[0-2]|0?[1-9])-(?:0?[1-9]|[1-2]\\d|30|31)) ((?:[01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d)$"
	regCarNo       = "^[京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼使领][A-HJ-NP-Z](?:(?:[A-HJ-NP-Z0-9]{4}[A-HJ-NP-Z0-9挂学警港澳])|(?:(?:\\d{5}[A-HJK])|(?:[A-HJK][A-HJ-NP-Z0-9][0-9]{4})))$"
	regTelePhone   = "^(?:(?:\\d{3}-)?\\d{8}|^(?:\\d{4}-)?\\d{7,8})(?:-\\d+)?$"
	regMobilePhone = "^(?:(?:\\+|00)86)?1[3-9]\\d{9}$"
	regIdCard      = "^[1-9]\\d{5}(?:18|19|20)\\d{2}(?:0[1-9]|10|11|12)(?:0[1-9]|[1-2]\\d|30|31)\\d{3}[0-9Xx]$"
	regDomain      = "^[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(?:\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+$"
	regUrl         = "^(?:https?:\\/\\/)?[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(?:\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+$"
	regEmail       = "^(?:https?:\\/\\/)?[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(?:\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+$"
	regIpv4        = "^(?:https?:\\/\\/)?[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(?:\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+$"
	regMAC         = "^(?:(?:[a-f0-9A-F]{2}:){5}|(?:[a-f0-9A-F]{2}-){5})[a-f0-9A-F]{2}$"
	regVersion     = "^\\d+(?:\\.\\d+){2}$"
)

func IsLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}
func IsUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func ToBytes(str string) []byte {
	return []byte(str)
}

func At(str string, index int) byte {
	if len(str) >= index {
		panic("index out of range")
	}
	return str[index]
}

func IsEmpty(str string) bool {
	return str == "" || len(str) == 0
}

func DefaultIfEmpty(str, defaultStr string) string {
	if IsEmpty(str) {
		return defaultStr
	} else {
		return str
	}
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}

func IsAnyEmpty(str ...string) bool {
	if len(str) == 0 {
		return true
	} else {
		for _, s := range str {
			if IsEmpty(s) {
				return true
			}
		}
		return false
	}
}

func IsNoneEmpty(str ...string) bool {
	return !IsAnyEmpty(str...)
}

func IsBlank(str string) bool {
	if IsEmpty(str) {
		return true
	} else {
		for _, s := range str {
			if s != whiteSpace {
				return false
			}
		}
	}
	return true
}

func DefaultIfBlank(str, defaultStr string) string {
	if IsBlank(str) {
		return defaultStr
	} else {
		return str
	}
}

func IsNotBlank(str string) bool {
	return !IsBlank(str)
}

func IsAnyBlank(str ...string) bool {
	if len(str) == 0 {
		return true
	}

	for _, s := range str {
		if IsBlank(s) {
			return true
		}
	}
	return false
}

func IsNoneBlank(str ...string) bool {
	return !IsAnyBlank(str...)
}

func MatchString(str, pattern string) bool {
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		return false
	}
	return matched
}

func FirstCharToUpper(str string) string {
	bytes := ToBytes(str)
	first := bytes[0]
	if first >= 'a' && first <= 'z' {
		bytes[0] = first - 32
		return string(bytes)
	} else {
		return str
	}
}

func FirstCharToLower(str string) string {
	bytes := ToBytes(str)
	first := bytes[0]
	if first >= 'A' && first <= 'Z' {
		bytes[0] = first + 32
		return string(bytes)
	} else {
		return str
	}
}

func ToLowerCamel(str string) string {

	splits := strings.Split(strings.Trim(str, "_"), "_")
	builder := strings.Builder{}

	for i, ele := range splits {
		if IsBlank(ele) {
			continue
		}

		var s = ToBytes(ele)
		if i == 0 {
			if IsUpper(s[0]) {
				s[0] += 32
			}
		} else {
			if IsLower(s[0]) {
				s[0] -= 32
			}
		}
		builder.Write(s)
	}
	return builder.String()
}

func ToUpperCamel(str string) string {
	splits := strings.Split(str, "_")
	builder := strings.Builder{}

	for _, ele := range splits {
		if IsBlank(ele) {
			continue
		}
		var s = ToBytes(ele)
		if IsLower(s[0]) {
			s[0] -= 32
		}
		builder.Write(s)
	}
	return builder.String()
}

func FillSlice(length int, fillStr string) []string {
	var slice []string
	for i := 0; i < length; i++ {
		slice = append(slice, fillStr)
	}
	return slice
}

//  dataMap = "a":"acorn" "b":"22"  str = "hello my name is {a},age is {b}"  => "hello my name is acorn,age is 22"
func FormatParam(str string, dataMap map[string]string) string {

	for k, v := range dataMap {
		str = strings.Replace(str, "{"+k+"}", v, 1)
	}
	return str
}

func Format(str, placeHolder string, args ...any) string {

	if IsBlank(str) || IsBlank(placeHolder) || args == nil || len(args) == 0 {
		return str
	}

	placeHolderLen := len(placeHolder)

	delimIndex := -1
	builder := strings.Builder{}

	for _, arg := range args {
		delimIndex = strings.Index(str, placeHolder)
		if delimIndex == -1 {
			break
		}
		builder.WriteString(str[:delimIndex])
		switch arg.(type) {
		case int:
			builder.WriteString(strconv.Itoa(arg.(int)))
		case float64:
			float := strconv.FormatFloat(arg.(float64), 'f', -1, 64)
			builder.WriteString(float)
		case string:
			builder.WriteString(arg.(string))
		case time.Time:
			t := arg.(time.Time)
			builder.WriteString(FormatTime(t))
		default:
			builder.WriteString(fmt.Sprintf("%v", arg))
		}
		delimIndex = delimIndex + placeHolderLen
		str = str[delimIndex:]
	}
	builder.WriteString(str)
	return builder.String()
}

func FormatSQL(sql string, params ...any) string {
	if IsBlank(sql) || params == nil || len(params) == 0 {
		return sql
	}

	delimIndex := -1
	builder := strings.Builder{}

	for _, arg := range params {
		delimIndex = strings.Index(sql, "?")
		if delimIndex == -1 {
			break
		}
		builder.WriteString(sql[:delimIndex])

		switch arg.(type) {
		case int:
			builder.WriteString(strconv.Itoa(arg.(int)))
		case float64:
			float := strconv.FormatFloat(arg.(float64), 'f', -1, 64)
			builder.WriteString(float)
		case string:
			builder.WriteString("'")
			builder.WriteString(arg.(string))
			builder.WriteString("'")
		case time.Time:
			builder.WriteString("'")
			t := arg.(time.Time)
			builder.WriteString(FormatTime(t))
			builder.WriteString("'")
		}

		delimIndex = delimIndex + 1
		sql = sql[delimIndex:]
	}
	builder.WriteString(sql)
	return builder.String()

}

func FormatTime(t time.Time) string {
	return t.Format(formatTime)
}

func FormatDistance(distance float64) string {

	if distance > 1000 {
		f := distance / 1000
		return strconv.FormatFloat(f, 'f', 2, 64) + "km"
	}
	return strconv.FormatFloat(distance, 'f', 2, 64) + "m"
}

func GetRandom() string {
	return GetRandomN(32)
}

func GetRandomN(count int) string {
	arr := []byte{}
	seedBytes := ToBytes(seedStr)

	length := len(seedStr)
	for i := 0; i < count; i++ {
		intn := rand.Intn(length)
		arr = append(arr, seedBytes[intn])
	}
	return string(arr)
}

func GetRandPrefix(prefix string, count int) string {
	length := count - len(prefix)
	return prefix + GetRandomN(length)
}

func GetCaptcha() string {
	return GetCaptchaN(4)
}

func GetCaptchaN(n int) string {
	return GetRandomN(n)
}

func IsDigit(str string) bool {
	return MatchString(str, regDigit)
}

func IsTime(str string) bool {
	return MatchString(str, regTime)
}

func IsDateTime(str string) bool {
	return MatchString(str, regDateTime)
}

func IsCarNo(str string) bool {
	return MatchString(str, regCarNo)
}

func IsTelePhone(str string) bool {
	return MatchString(str, regTelePhone)
}

func IsMobilePhone(str string) bool {
	return MatchString(str, regMobilePhone)
}

func IsIdCard(str string) bool {
	return MatchString(str, regIdCard)
}

func IsDomain(str string) bool {
	return MatchString(str, regDomain)
}

func IsUrl(str string) bool {
	return MatchString(str, regUrl)
}

func IsEmail(str string) bool {
	return MatchString(str, regEmail)
}
func IsIpV4(str string) bool {
	return MatchString(str, regIpv4)
}

func IsMAC(str string) bool {
	return MatchString(str, regMAC)
}

func IsVersionNo(str string) bool {
	return MatchString(str, regVersion)
}
