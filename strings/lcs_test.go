package strings

import "testing"

func TestLongestSubSeqLenDPReducedSpace(t *testing.T) {
	data := []struct {
		s1, s2 string
		out    int
	}{
		{"ABAZDC", "BACBAD", 4},
		{"AGGTAB", "GXTXAYB", 4},
		{"aaaa", "aa", 2},
		{"ABBA", "ABCABA", 4},
	}

	for i, d := range data {
		res := LongestSubSeqLenDPReducedSpace(d.s1, d.s2)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}
func TestLongestSubSeqDP(t *testing.T) {
	data := []struct {
		s1, s2, out string
	}{
		{"ABAZDC", "BACBAD", "ABAD"},
		{"AGGTAB", "GXTXAYB", "GTAB"},
		{"aaaa", "aa", "aa"},
		{"ABBA", "ABCABA", "ABBA"},
	}

	for i, d := range data {
		res := LongestSubSeqDP(d.s1, d.s2)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}

func TestLongestSubSeqLenDP(t *testing.T) {
	data := []struct {
		s1, s2 string
		out    int
	}{
		{"ABAZDC", "BACBAD", 4},
		{"AGGTAB", "GXTXAYB", 4},
		{"aaaa", "aa", 2},
		{"ABBA", "ABCABA", 4},
	}

	for i, d := range data {
		res := LongestSubSeqLenDP(d.s1, d.s2)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}

func TestLongestSubSeqLenWithMemoization(t *testing.T) {
	data := []struct {
		s1, s2 string
		out    int
	}{
		{"ABAZDC", "BACBAD", 4},
		{"AGGTAB", "GXTXAYB", 4},
		{"aaaa", "aa", 2},
		{"ABBA", "ABCABA", 4},
	}

	for i, d := range data {
		res := LongestSubSeqLenWithMemoization(d.s1, d.s2)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}

func TestLongestSubSeqLen(t *testing.T) {
	data := []struct {
		s1, s2 string
		out    int
	}{
		{"ABAZDC", "BACBAD", 4},
		{"AGGTAB", "GXTXAYB", 4},
		{"aaaa", "aa", 2},
		{"ABBA", "ABCABA", 4},
	}

	for i, d := range data {
		res := LongestSubSeqLen(d.s1, d.s2, 0, 0)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}

func TestSubSeq(t *testing.T) {
	data := []struct {
		s1, pat string
		out     bool
	}{
		{"ABAZDC", "ABAD", true},
		{"AGGTAB", "WTAB", false},
		{"aaaa", "aac", false},
		{"ABCABA", "ABBA", true},
		{"anematode knowledge", "nano", true},
	}

	for i, d := range data {
		res := SubSeq(d.pat, d.s1)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}

func TestLongestSubSeqBF(t *testing.T) {
	data := []struct {
		s1, s2, out string
	}{
		{"ABAZDC", "BACBAD", "ABAD"},
		{"AGGTAB", "GXTXAYB", "GTAB"},
		{"aaaa", "aa", "aa"},
		{"ABBA", "ABCABA", "ABBA"},
	}

	for i, d := range data {
		res := LongestSubSeqBF(d.s1, d.s2)
		if res != d.out {
			t.Errorf("case [%v] expected %v, got %v", i, d.out, res)
		}
	}
}
