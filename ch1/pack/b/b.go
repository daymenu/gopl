package b

import (
	"fmt"

	"github.com/gopl/ch1/pack/c"
)

var mySex = Sex
var Name = "b"

var Sex = 1

func init() {
	fmt.Println("init b:", c.Name)
}
