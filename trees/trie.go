package tree

// Trie for case insensitive ascii chars
// panics on insert and find if chars out of range
func NewTrieForCaseInsensitveAlphabet26() *Trie {
	return &Trie{root: &trieNode{
		children: make([]*trieNode, 26), isEndOfWord: false},
		alphabetSize: 26,
		charToAbcID: func(ch byte) int {
			if ch > 64 && ch < 91 {
				ch += 32
			} else if !(ch > 96 && ch < 123) {
				panic("wrong alphabet char " + string(ch))
			}

			return int(ch - 'a')
		},
	}
}

func NewTrie(alphabetSize int, charToAbcID func(byte) int) *Trie {
	return &Trie{root: &trieNode{
		children: make([]*trieNode, alphabetSize), isEndOfWord: false},
		alphabetSize: alphabetSize,
		charToAbcID:  charToAbcID,
	}
}

// Trie with lexicographic sorting
type Trie struct {
	root         *trieNode
	alphabetSize int
	charToAbcID  func(ch byte) int
}

type trieNode struct {
	children    []*trieNode
	val         interface{}
	isEndOfWord bool
}

func (t *Trie) newNode() *trieNode {
	return &trieNode{
		children:    make([]*trieNode, t.alphabetSize),
		isEndOfWord: false}
}

type TrieWalker struct {
	t   *Trie
	n   *trieNode
	val interface{}
}

func (w *TrieWalker) Next(ch byte) *TrieWalker {
	node := w.n.children[w.t.charToAbcID(ch)]
	if node == nil {
		return nil
	}
	w.n = node
	return w
}

func (w *TrieWalker) IsWord() bool {
	return w.n.isEndOfWord
}

func (w *TrieWalker) NodeVal() interface{} {
	return w.n.val
}

func (t *Trie) NewCharWalker() *TrieWalker {
	return &TrieWalker{t, t.root, nil}
}

// for ascii chars
func (t *Trie) Insert(str string, val interface{}) {
	node := t.root
	for i := 0; i < len(str); i++ {
		ch := t.charToAbcID(str[i])
		if node.children[ch] == nil {
			node.children[ch] = t.newNode()
		}
		node = node.children[ch]
	}
	node.isEndOfWord = true
	node.val = val
}

// for ascii chars
func (t *Trie) Find(str string) (interface{}, bool) {
	node := t.root
	for i := 0; i < len(str); i++ {
		node = node.children[t.charToAbcID(str[i])]
		if node == nil {
			return nil, false
		}
	}

	if node.isEndOfWord {
		return node.val, true
	}
	return nil, false
}

func (t *Trie) String() []string {
	var out []string
	var str []byte
	var walk func(*trieNode, int, []byte)
	walk = func(node *trieNode, ch int, str []byte) {
		if node == nil {
			return
		}
		str = append(str, byte(ch)+'a')
		if node.isEndOfWord {
			out = append(out, string(str))
		}
		for ch, n := range node.children {
			walk(n, ch, str)
		}
	}
	for ch, n := range t.root.children {
		walk(n, ch, str)
		str = str[:0]
	}
	return out
}
