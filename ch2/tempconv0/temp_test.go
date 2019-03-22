package tempconv0

import "testing"

func TestCToF(t *testing.T) {
	if CToF(Celsius(BoilingC)) != Fahrenheit(212.200000) {
		t.Errorf("%s = %s ", CToF(Celsius(BoilingC)), Fahrenheit(0))
	}
}
