package strings

import (
	"testing"

	"github.com/mmirolim/algos/checks"
)

func TestCheckParenthesesBalanced03(t *testing.T) {
	for i, d := range []struct {
		str string
		pos int
	}{
		{"(())", -1},
		{"((())())()", -1},
		{")()(", 0},
		{"())", 2},
	} {
		pos := CheckParenthesesBalanced03(d.str)
		if pos != d.pos {
			t.Errorf("case [%d] expected %+v, got %+v", i, d.pos, pos)
		}

	}
}

func TestCheckParenthesesBalanced02(t *testing.T) {
	for i, d := range []struct {
		str string
		pos int
	}{
		{"(())", -1},
		{"((())())()", -1},
		{")()(", 0},
		{"())", 2},
	} {
		pos := CheckParenthesesBalanced02(d.str)
		if pos != d.pos {
			t.Errorf("case [%d] expected %+v, got %+v", i, d.pos, pos)
		}

	}
}

func TestCheckParenthesesBalanced01(t *testing.T) {
	for i, d := range []struct {
		str string
		pos int
	}{
		{"(())", -1},
		{"((())())()", -1},
		{")()(", 0},
		{"())", 2},
	} {
		pos := CheckParenthesesBalanced01(d.str)
		if pos != d.pos {
			t.Errorf("case [%d] expected %+v, got %+v", i, d.pos, pos)
		}

	}
}

func TestRabinKarpSubStringSearch(t *testing.T) {
	prime := 619
	for i, data := range []struct {
		Str, Pat string
		expected []int
	}{
		{"", "", []int{0}},
		{"AABAACAADAABAABA", "AABA", []int{0, 9, 12}},
		{"ALKJFDF", "ZAB", []int{-1}},
	} {
		out := RabinKarpSubStringSearch(data.Pat, data.Str, prime)
		if len(out) != len(data.expected) {
			t.Errorf("case [%d] expected %v, got %v", i, data.expected, out)
			continue
		}
		for _, j := range out {
			if j == -1 && data.expected[0] == -1 {
				break
			}

			subStr := string(data.Str[j : j+len(data.Pat)])
			if subStr != data.Pat {
				t.Errorf("case [%d] expected %+v, got %v \n",
					i, data.expected, out) // output for debug
				break
			}
		}
	}

}

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
func TestSearchStrCharsInJournal(t *testing.T) {
	for i, data := range []struct {
		Str, Pat string
		expected bool
	}{
		{"ABDOECBOKABABKC", "ABC", true},
		{"this is a test string", "tist", true},
		{"geeksforgeeks", "work", false},
	} {
		out := SearchStrCharsInJournal(data.Pat, data.Str)

		if out != data.expected {
			t.Errorf("case [%d] expected %+v, got %v \n",
				i, data.expected, out) // output for debug
		}
	}

}

func TestReverseWordsInSentence(t *testing.T) {
	in := "My name is Chris"
	out := "Chris is name My"
	res := ReverseWordsInSentence(in)
	if res != out {
		t.Errorf("expected %+v, got %+v", out, res)
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

func TestCryptKickerDecodeString01(t *testing.T) {
	for i, d := range []struct {
		words    []string
		str, out string
	}{{[]string{
		"and",
		"puff",
		"dick",
		"jane",
		"yertle",
		"spot",
	},

		"bjvg xsb hxsn xsb qymm xsb rqat xsb pnetfn", "dick and jane and puff and spot and yertle"},

		{[]string{
			"and",
			"dick",
			"jane",
			"puff",
			"spot",
			"yertle",
		}, "bjvg xsb hxsn xsb qymm xsb rqat xsb pnetfn", "dick and jane and puff and spot and yertle"},
		{[]string{
			"and",
			"dick",
			"jane",
			"puff",
			"spot",
			"yertle",
		}, "xxxx yyy zzzz www yyyy aaa bbbb ccc dddddd", "**** *** **** *** **** *** **** *** ******"},
	} {
		out := CryptKickerDecodeString01(d.words, d.str)
		if out != d.out {
			t.Errorf("case [%d] expected %+v, got %v \n",
				i, d.out, out) // output for debug
		}
	}

}
func TestCryptKickerDecodeStringRecur(t *testing.T) {
	for i, d := range []struct {
		words    []string
		str, out string
	}{
		{[]string{
			"and",
			"puff",
			"dick",
			"jane",
			"yertle",
			"spot",
		},

			"bjvg xsb hxsn xsb qymm xsb rqat xsb pnetfn", "dick and jane and puff and spot and yertle"},
		{[]string{
			"and",
			"dick",
			"jane",
			"puff",
			"spot",
			"yertle",
		}, "bjvg xsb hxsn xsb qymm xsb rqat xsb pnetfn", "dick and jane and puff and spot and yertle"},
		{[]string{
			"and",
			"dick",
			"jane",
			"puff",
			"spot",
			"yertle",
		}, "xxxx yyy zzzz www yyyy aaa bbbb ccc dddddd", "**** *** **** *** **** *** **** *** ******"},
	} {
		out := CryptKickerDecodeStringRecur(d.words, d.str)
		if out != d.out {
			t.Errorf("case [%d] expected %+v, got %v \n",
				i, d.out, out) // output for debug
		}
	}

}

func TestWhereIsWaldorfFindString(t *testing.T) {
	data := []struct {
		grid  []string
		words []string
		pos   []int
	}{
		{
			[]string{"abcDEFGhigg",
				"hEbkWalDork",
				"FtyAwaldORm",
				"FtsimrLqsrc",
				"byoArBeDeyv",
				"Klcbqwikomk",
				"strEBGadhrb",
				"yUiqlxcnBjf"},
			[]string{
				"Waldorf",
				"Bambi",
				"Betty",
				"Dagbert"},
			[]int{2, 5, 2, 3, 1, 2, 7, 8},
		},
	}

	for i, d := range data {
		out := WhereIsWaldorfFindString(d.grid, d.words)
		for j := range d.words {
			if d.pos[j*2] != out[j*2] || d.pos[j*2+1] != out[j*2+1] {
				t.Errorf("case [%v] expected row %v cow %v got row %v col %v for word %v", i, d.pos[j*2], d.pos[j*2+1], out[j*2], out[j*2+1], d.words[j])
			}
		}
	}
}

func TestWhereIsWaldorfFindStringUsingTrie(t *testing.T) {
	data := []struct {
		grid  []string
		words []string
		pos   []int
	}{
		{
			[]string{"abcDEFGhigg",
				"hEbkWalDork",
				"FtyAwaldORm",
				"FtsimrLqsrc",
				"byoArBeDeyv",
				"Klcbqwikomk",
				"strEBGadhrb",
				"yUiqlxcnBjf"},
			[]string{
				"Waldorf",
				"Bambi",
				"Betty",
				"Dagbert",
			},
			[]int{2, 5, 2, 3, 1, 2, 7, 8},
		},
	}

	for i, d := range data {
		out := WhereIsWaldorfFindStringUsingTrie(d.grid, d.words)
		for j := range d.words {
			if d.pos[j*2] != out[j*2] || d.pos[j*2+1] != out[j*2+1] {
				t.Errorf("case [%v] expected row %v cow %v got row %v col %v for word %v", i, d.pos[j*2], d.pos[j*2+1], out[j*2], out[j*2+1], d.words[j])
			}
		}
	}
}

func TestCryptKickerIIDecodeString(t *testing.T) {
	for i, d := range []struct {
		knownLine string
		lines     []string
		out       []string
	}{
		{"the quick brown fox jumps over the lazy dog",
			[]string{
				"vtz ud xnm xugm itr pyy jttk gmv xt otgm xt xnm puk ti xnm fprxq",
				"xnm ceuob lrtzv ita hegfd tsmr xnm ypwq ktj",
				"frtjrpgguvj otvxmdxd prm iev prmvx xnmq",
			},
			[]string{
				"now is the time for all good men to come to the aid of the party",
				"the quick brown fox jumps over the lazy dog",
				"programming contests are fun arent they",
			},
		},
	} {
		out := CryptKickerIIDecodeString(d.knownLine, d.lines)
		if len(out) != len(d.out) {
			t.Errorf("case [%d] expected %+v, got %v \n",
				i, len(d.out), len(out)) // output for debug
		}
		for j := range out {
			if out[j] != d.out[j] {
				t.Errorf("case [%d] expected %+v, got %v \n",
					i, d.out, out) // output for debug
			}
		}
	}

}
func TestPermutate(t *testing.T) {
	data := []struct {
		str   string
		perms []string
	}{
		{"AB", []string{"AB", "BA"}},
		{"123", []string{"123", "132", "213", "231", "321", "312"}},
	}
	for _, d := range data {
		out := Permutate(d.str)
		checks.AssertEq(t, d.perms, out)
	}
}

func BenchmarkWhereIsWaldorfFindString(b *testing.B) {
	grid := []string{
		"abcDEFGhiggabcDEFGhiggabcDEFGhiggabcDEFGhiggabcDEFGhigg",
		"hEbkWalDorkhEbkWalDorkhEbkWalDorkhEbkWalDorkhEbkWalDork",
		"FtyAwaldORmFtyAwaldORmFtyAwaldORmFtyAwaldORmFtyAwaldORm",
		"FtsimrLqsrcFtsimrLqsrcFtsimrLqsrcFtsimrLqsrcFtsimrLqsrc",
		"byoArBeDeyvbyoArBeDeyvbyoArBeDeyvbyoArBeDeyvbyoArBeDeyv",
		"KlcbqwikomkKlcbqwikomkKlcbqwikomkKlcbqwikomkKlcbqwikomk",
		"strEBGadhrbstrEBGadhrbstrEBGadhrbstrEBGadhrbstrEBGadhrb",
		"yUiqlxcnBjfyUiqlxcnBjfyUiqlxcnBjfyUiqlxcnBjfyUiqlxcnBjf"}

	words := []string{
		"Waldorf",
		"Bambi",
		"Betty",
		"Dagbert",
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = WhereIsWaldorfFindString(grid, words)
	}
}

func BenchmarkWhereIsWaldorfFindStringUsingTrie(b *testing.B) {
	grid := []string{
		"abcDEFGhiggabcDEFGhiggabcDEFGhiggabcDEFGhiggabcDEFGhigg",
		"hEbkWalDorkhEbkWalDorkhEbkWalDorkhEbkWalDorkhEbkWalDork",
		"FtyAwaldORmFtyAwaldORmFtyAwaldORmFtyAwaldORmFtyAwaldORm",
		"FtsimrLqsrcFtsimrLqsrcFtsimrLqsrcFtsimrLqsrcFtsimrLqsrc",
		"byoArBeDeyvbyoArBeDeyvbyoArBeDeyvbyoArBeDeyvbyoArBeDeyv",
		"KlcbqwikomkKlcbqwikomkKlcbqwikomkKlcbqwikomkKlcbqwikomk",
		"strEBGadhrbstrEBGadhrbstrEBGadhrbstrEBGadhrbstrEBGadhrb",
		"yUiqlxcnBjfyUiqlxcnBjfyUiqlxcnBjfyUiqlxcnBjfyUiqlxcnBjf"}

	words := []string{
		"Waldorf",
		"Bambi",
		"Betty",
		"Dagbert",
	}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = WhereIsWaldorfFindStringUsingTrie(grid, words)
	}
}
