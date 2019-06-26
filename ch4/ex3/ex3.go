package ex3

// 数组长度
const RLen = 10

// 反转数组
func Reverse(arr *[RLen]int) {
	for i := 0; i < RLen/2; i++ {
		arr[i], arr[RLen-i-1] = arr[RLen-i-1], arr[i]
	}
}
