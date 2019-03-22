package basename2

import "strings"

func BaseName(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot != -1 {
		s = s[:dot]
	}
	return s
}
