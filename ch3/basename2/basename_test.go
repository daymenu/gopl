package basename2

import "testing"

func TestBaseName(t *testing.T) {
	paths := []struct {
		path string
		want string
	}{
		{"D:/dd/ff.gif", "ff"},
		{"D:/dd/ff", "ff"},
		{"D:/dd/", ""},
		{"", ""},
		{"/", ""},
	}
	for _, v := range paths {
		if baseName := BaseName(v.path); baseName != v.want {
			t.Errorf("BaseName(%s)=%s,want %s", v.path, baseName, v.want)
		}
	}
}
