package ex6

// 将相邻的多个空格合并为一个
func MergeSpace(b []byte) []byte {
	space := byte(' ')
	front := 0
	for i := 1; i < len(b); i++ {
		if b[i] == space && b[front] == space {
			continue
		}
		front++
		if i-front > 0 {
			b[front] = b[i]
		}
	}
	return b[:front+1]
}
