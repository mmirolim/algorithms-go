package tree

import (
	"testing"
)

func TestNewTrieForCaseInsensitveAlphabet26(t *testing.T) {
	trie := NewTrieForCaseInsensitveAlphabet26()
	if trie.alphabetSize != 26 {
		t.Errorf("expected %v, got %v", 26, trie.alphabetSize)
	} else if trie.charToAbcID == nil {
		t.Errorf("expected %v, got %v", "nil", "not nil")
	} else if len(trie.root.children) != 26 {
		t.Errorf("expected %v, got %v", 26, len(trie.root.children))
	}

	charsIn := []byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
		'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}
	charsOut := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
	}
	for i := range charsIn {
		ch := trie.charToAbcID(charsIn[i])
		if ch != charsOut[i] {
			t.Errorf("expected %v, got %v", charsOut[i], ch)
		}
	}
}
func TestTrieInsert(t *testing.T) {
	trie := NewTrieForCaseInsensitveAlphabet26()
	input := []string{"A", "to", "TEA", "ted", "ten", "inn"}
	for i, str := range input {
		trie.Insert(str, i)
	}
	fn := trie.charToAbcID
	n := trie.root
	nodeA := n.children[fn('A')]
	nodeTEA := n.children[fn('T')].
		children[fn('E')].children[fn('A')]
	nodeted := n.children[fn('t')].
		children[fn('e')].children[fn('d')]
	nodeinn := n.children[fn('i')].
		children[fn('n')].children[fn('n')]
	for i, d := range []struct {
		node        *trieNode
		isEndOfWord bool
		val         int
	}{
		{nodeA, true, 0}, {nodeTEA, true, 2}, {nodeted, true, 3}, {nodeinn, true, 5},
	} {
		if d.node.isEndOfWord != d.isEndOfWord || d.node.val != d.val {
			t.Errorf("case [%v] expected isEndOfWord %+v and val %+v, got %+v %+v",
				i, d.node.isEndOfWord, d.node.val, d.isEndOfWord, d.val)
		}
	}
}

func TestTrieFind(t *testing.T) {
	trie := NewTrieForCaseInsensitveAlphabet26()
	input := []string{"A", "to", "TEA", "ted", "ten", "inn"}
	for i, str := range input {
		trie.Insert(str, i)
	}
	for i, d := range []struct {
		key   string
		found bool
		val   interface{}
	}{
		{"X", false, nil}, {"A", true, 0}, {"inn", true, 5}, {"hello", false, nil},
	} {
		val, ok := trie.Find(d.key)
		if ok != d.found || val != d.val {
			t.Errorf("case [%v] expected found %+v and val %+v, got %+v %+v",
				i, d.found, d.val, ok, val)
		}
	}
}

func TestTrieString(t *testing.T) {
	trie := NewTrieForCaseInsensitveAlphabet26()
	input := []string{"A", "to", "TEA", "ted", "ten", "inn"}
	expected := []string{"a", "inn", "tea", "ted", "ten", "to"}
	for i, str := range input {
		trie.Insert(str, i)
	}
	out := trie.String()
	if len(expected) != len(out) {
		t.Errorf("expected %v, got %v", expected, out)
	}

	for i := range out {
		if expected[i] != out[i] {
			t.Errorf("expected %v, got %v", expected, out)
		}
	}
}
