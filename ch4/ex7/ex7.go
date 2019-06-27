package ex7

func reverse(b []byte) []byte {
	bLen := len(b)
	halfLen := len(b) / 2
	for i := 0; i < halfLen; i++ {
		b[i], b[bLen-i] = b[bLen-i], b[i]
	}
	return b
}
