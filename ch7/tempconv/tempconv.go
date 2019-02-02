package tempconv

import (
	"flag"
	"fmt"
	"strings"
)

type Celsius float64
type Fahrenheit float64
type Kelvins float64
type celsiusFlag struct{ Celsius }

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * (5.0 / 9.0))
}
func KToC(k Kelvins) Celsius {
	return Celsius(k + 273.15)
}
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch strings.ToUpper(unit) {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvins(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f °C", c)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
