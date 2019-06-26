package mapequal

func Equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for xk, xv := range x {
		if yv, ok := y[xk]; !ok || xv != yv {
			return false
		}
	}
	return true
}
