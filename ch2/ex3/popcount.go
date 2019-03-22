package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	bn := 0
	for i := 0; i < 8; i++ {
		bn += int(pc[byte(x>>uint((i*8)))])
	}
	return bn
}

func FPopCount(x uint64) int {
	num := 0
	var i uint8
	for i = 0; i < 8; i++ {
		num += int(pc[byte(x>>(i*8))])
	}
	return num
}
