package comma

import "testing"

type Money struct {
	money string
	want  string
}

func TestComma(t *testing.T) {
	moneys := getMoneys()
	for _, v := range moneys {
		if c := Comma(v.money); c != v.want {
			t.Errorf("Comma(%s)=%s,want:%s", v.money, c, v.want)
		}
	}
}
func TestRecursiveComma(t *testing.T) {
	moneys := getMoneys()
	for _, v := range moneys {
		if c := RecursiveComma(v.money); c != v.want {
			t.Errorf("Comma(%s)=%s,want:%s", v.money, c, v.want)
		}
	}
}
func getMoneys() []Money {
	return []Money{
		{"123456789", "123,456,789"},
		{"123", "123"},
		{"1234", "1,234"},
		{"-123456789.6432", "-123,456,789.6432"},
		{"+89.6432", "+89.6432"},
		{"100.6432", "100.6432"},
	}
}
