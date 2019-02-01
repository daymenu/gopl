package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type celsiusFlag struct{ Celsius }

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * (5.0 / 9.0))
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Scanf(s, "%f%s", &value, &unit)
	switch unit {
	case "c", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
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
