package ex6

import (
	"testing"
)

type TestData struct {
	b []byte
	r string
}

func TestMergeSpace(t *testing.T) {
	tds := []TestData{
		// {[]byte("i am a    boy"), "i am a boy"},
		// {[]byte("     am a    boy"), " am a boy"},
		// {[]byte("     "), " "},
		{[]byte("eee     "), "eee "},
	}
	for _, td := range tds {
		r := MergeSpace(td.b)
		if string(r) != td.r {
			t.Errorf("%0s != %0s", r, td.r)
		}
	}

}
