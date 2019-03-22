// 给钱加逗号
// 思路从后往前数，每逢3个加逗号，最后一位不需要，小于3位不需要
package comma

import (
	"bytes"
	"strings"
)

func Comma(s string) string {
	b := new(bytes.Buffer)
	fint := s
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		b.WriteString(s[:1])
		fint = s[1:]
	}
	decimal := ""
	if i := strings.Index(fint, "."); i != -1 {
		decimal, fint = fint[i:], fint[:i]
	}
	n := len(fint)
	if n < 3 {
		return s
	}
	pre := n % 3
	if pre != 0 {
		b.WriteString(fint[:pre] + ",")
		fint = fint[pre:]
	}
	for i := 3; i <= len(fint); i += 3 {
		if i == len(fint) {
			s1 := fint[i-3 : i]
			b.WriteString(s1)
		} else {
			s2 := fint[i-3:i] + ","
			b.WriteString(s2)
		}
	}
	b.WriteString(decimal)
	return b.String()
}

func RecursiveComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return RecursiveComma(s[:n-3]) + "," + s[n-3:]
}
