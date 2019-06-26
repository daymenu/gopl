package ex5

// ClearSameChar 消除相同字符操作
func ClearSameChar(s []string) []string {
	if len(s) == 0 {
		return s
	}
	front := 0
	for i := range s {
		if s[front] != s[i] {
			front++
			s[front] = s[i]
		}
	}
	return s[:front+1]
}
