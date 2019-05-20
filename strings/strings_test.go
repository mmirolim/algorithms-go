package strings

import (
	"testing"
)

func TestSearchMinWindow(t *testing.T) {
	for i, data := range []struct {
		Str, Pat, expected string
	}{
		// TODO add test when pat is missing in string
		{"ABDOECBOKABABKC", "ABC", "ABKC"},
		{"this is a test string", "tist", "t stri"},
		{"geeksforgeeks", "ork", "ksfor"},
	} {
		index, minlen := searchMinWindow(data.Pat, data.Str)

		substr := data.Str[index : index+minlen]
		if substr != data.expected {
			t.Errorf("case [%d] got substring %+v expected %+v\n",
				i, substr, data.expected) // output for debug
		}
	}

}

func BenchmarkSearchMinWindow(b *testing.B) {
	pat := "tist"
	str := "this is a test string"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = searchMinWindow(pat, str)
	}
}
