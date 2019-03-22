package myappend

func MyAppend(s []rune, a ...rune) []rune {
	var z []rune
	zlen := len(s) + len(a)

	if zlen <= cap(z) {
		z = s[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(s) {
			zcap = 2 * zlen
		}
		z = make([]rune, zlen, zcap)
		copy(z, s)
	}

	copy(z[zlen:], a)
	return z
}
