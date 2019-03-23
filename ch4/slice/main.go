package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "Augst", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q1 := months[1:4]
	Q2 := months[4:7]
	fmt.Println(Q1, len(Q1), cap(Q1))
	fmt.Println(Q2, len(Q2), cap(Q2))
}
