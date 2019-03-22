package ex4

func PopCount(x uint64) int {
	bn := 0
	var i uint
	for i = 0; i < 64; i++ {
		bn += int((x & (1 << i)) >> i)
	}
	return bn
}
