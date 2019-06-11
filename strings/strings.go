package strings

import (
	"strings"
)

type chIndex struct {
	b  byte
	id int
}

func newCharStack() *[]chIndex {
	return &[]chIndex{}
}

func push(ch byte, pos int, st *[]chIndex) {
	*st = append(*st, chIndex{ch, pos})
}
func pop(st *[]chIndex) chIndex {
	sl := *st
	ch := sl[len(sl)-1]
	*st = sl[0 : len(sl)-1]
	return ch
}

func peek(st *[]chIndex) chIndex {
	return (*st)[len(*st)-1]
}

func CheckParenthesesBalanced03(str string) int {
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			count++
		} else if str[i] == ')' {
			count--
		} else {
			panic("unexpected char" + string(str[i]))
		}

		if count < 0 {
			return i
		}
	}

	return -1
}

func CheckParenthesesBalanced02(str string) int {
	st := newCharStack()
	if str[0] == ')' {
		return 0
	}
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			push(str[i], i, st)
		} else if str[i] == ')' {
			if len(*st) > 0 {
				ch := pop(st)
				if ch.b != '(' {
					return i
				}
			} else {
				return i
			}
		} else {
			panic("unexpected char" + string(str[i]))
		}

	}

	return -1
}

func CheckParenthesesBalanced01(str string) int {
	st := newCharStack()
	if str[0] == ')' {
		return 0
	}
	for i := 0; i < len(str); i++ {
		if len(*st) > 0 {
			ch := peek(st)
			if ch.b+1 == str[i] {
				pop(st)
				continue
			}
		}
		push(str[i], i, st)

	}

	if len(*st) == 0 {
		return -1
	}
	return (*st)[0].id
}

// problem ADM 3.24
func SearchStrCharsInJournal(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	freq1 := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		freq1[s1[i]]++
	}

	for i := range s2 {
		k := s2[i]
		_, ok := freq1[k]
		if ok {
			freq1[k]--
			if freq1[k] == 0 {
				delete(freq1, k)
				if len(freq1) == 0 {
					return true
				}
			}
		}
	}

	return len(freq1) == 0
}

// problem ADM 3.26
func ReverseWordsInSentence(str string) string {
	out := make([]byte, len(str))
	l := 0
	lastId := len(str)
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			k := i - l
			for j := lastId - l; j < lastId; j++ {
				out[j] = str[k]
				k++
			}
			lastId = lastId - l - 1
			out[lastId] = ' '
			l = 0
		} else {
			l++
		}
	}
	if l > 0 {
		for j := 0; j < l; j++ {
			out[j] = str[len(str)-l+j]
		}
	}

	return string(out)
}

/*
Crypt Kicker
PC/UVa IDs: 110204/843,
Popularity: B, Success rate: low Level: 2
*/
func CryptKickerDecodeString01(words []string, text string) string {
	if len(words) == 0 || len(text) == 0 {
		return ""
	}
	out := make([]byte, len(text))
	for i := range text {
		if text[i] == ' ' {
			out[i] = ' '
		} else {
			out[i] = '*'
		}
	}

	dicTextWordToDicWord := map[string]string{}

	createDic := func(ws []string) map[string]struct{} {
		m := map[string]struct{}{}
		for i := range ws {
			m[ws[i]] = struct{}{}
		}
		return m
	}
	createDicByLen := func(dic map[string]struct{}) map[int][]string {
		m := map[int][]string{}
		for k := range dic {
			m[len(k)] = append(m[len(k)], k)
		}
		return m
	}

	textDic := createDic(strings.Split(text, " "))
	textDicByLen := createDicByLen(textDic)
	wordsDic := createDic(words)
	wordsDicByLen := createDicByLen(wordsDic)
	// check that dic big enough
	for k, v := range wordsDicByLen {
		if len(v) < len(textDicByLen[k]) {
			return string(out)
		}
	}
	setIDx := 0
	wordIDx := 0
	sets := make([][]int, len(textDicByLen))
	textWordsBySetIndex := map[int][]string{}
	i := 0
	for _, s := range textDicByLen {
		sets[i] = make([]int, len(s))
		for j := range sets[i] {
			sets[i][j] = -1
		}
		textWordsBySetIndex[i] = s
		i++
	}

	IsAllLettersOneToOne := func() bool {
		dicAbc := uint64(0)
		abc := [26]int{}
		for i := range abc {
			abc[i] = -1
		}
		for k, v := range dicTextWordToDicWord {
			for i := range k {
				if abc[k[i]-97] == -1 {
					if dicAbc&(1<<(v[i]-97)) != 0 {
						return false
					}
					abc[k[i]-97] = int(v[i])
					dicAbc |= (1 << (v[i] - 97))
				} else if abc[k[i]-97] != int(v[i]) {
					return false
				}
			}
		}
		return true
	}

	lastWordVariant := func() bool {
		out := sets[setIDx][wordIDx] >
			len(wordsDicByLen[len(textWordsBySetIndex[setIDx][0])])-1
		return out
	}

	check := func() bool {
		if lastWordVariant() {
			return false
		}

		dicTextWordToDicWord[textWordsBySetIndex[setIDx][wordIDx]] =
			wordsDicByLen[len(textWordsBySetIndex[setIDx][0])][sets[setIDx][wordIDx]]
		out := IsAllLettersOneToOne()
		return out
	}

	isFirstWordInSet := func() bool {
		return wordIDx == 0
	}
	regressToPrevSet := func() bool {
		setIDx--
		if setIDx < 0 {
			return false
		}
		wordIDx = len(sets[setIDx]) - 1
		return true
	}
	regress := func() bool {
		if isFirstWordInSet() {
			return regressToPrevSet()
		}
		wordIDx--
		return true
	}

	nextSet := func() {
		wordIDx = 0
		setIDx++

	}
	isLastIndexInSet := func() bool {
		return wordIDx >= len(sets[setIDx])-1
	}
	advance := func() {
		if isLastIndexInSet() {
			nextSet()
		} else {
			wordIDx++
		}
	}
	noSetsLeft := func() bool {
		return setIDx == len(sets)
	}
	tryWord := func() {
		sets[setIDx][wordIDx]++

	}
	removeWordMatch := func() {
		k := textDicByLen[len(textWordsBySetIndex[setIDx][0])][wordIDx]
		delete(dicTextWordToDicWord, k)
	}
	clearSetCell := func() {
		sets[setIDx][wordIDx] = -1
	}
	for len(dicTextWordToDicWord) != len(textDic) {
		tryWord()
		if check() {
			advance()
			if noSetsLeft() {
				break
			}
		} else {
			removeWordMatch()
			if lastWordVariant() {
				clearSetCell()
				if !regress() {
					removeWordMatch()
					break
				}
			}
		}
	}
	generateAbc := func() map[byte]byte {
		out := map[byte]byte{}
		for k, v := range dicTextWordToDicWord {
			for i := range k {
				out[k[i]] = v[i]
			}
		}
		return out
	}
	if len(dicTextWordToDicWord) == len(textDic) {
		abc := generateAbc()
		for i := range text {
			v, ok := abc[text[i]]
			if ok {
				out[i] = v
			} else {
				out[i] = text[i]
			}
		}
	}

	return string(out)
}

func Permutate(s string) []string {
	out := []string{}
	var permutate func(str []byte, l, r int)
	permutate = func(str []byte, l, r int) {
		if l == r {
			out = append(out, string(str))
		} else {
			for i := l; i <= r; i++ {
				str[l], str[i] = str[i], str[l]
				permutate(str, l+1, r)
			}
		}
	}
	permutate([]byte(s), 0, len(s)-1)
	return out

}

/*
Crypt Kicker
PC/UVa IDs: 110204/843,
Popularity: B, Success rate: low Level: 2
*/
func CryptKickerDecodeStringRecur(words []string, text string) string {
	out := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			out[i] = ' '
		} else {
			out[i] = '*'
		}
	}
	textWords := func() []string {
		var out []string
		m := make(map[string]bool)
		l := 0
		for i := 0; i < len(text); i++ {
			if text[i] == ' ' {
				m[text[i-l:i]] = true
				l = 0
			} else {
				l++
			}
		}
		if l > 0 {
			m[text[len(text)-1-l+1:]] = true
		}
		for k := range m {
			out = append(out, k)
		}
		return out
	}()

	abc := make([]byte, 27)
	abcInverse := make([]byte, 27)
	baseNum := byte(96)

	setChar := func(a, x byte) bool {
		aInAbc := abc[a-baseNum]
		xInAbcInverse := abcInverse[x-baseNum]
		// not yet set
		if aInAbc == 0 && xInAbcInverse == 0 {
			abc[a-baseNum] = x
			abcInverse[x-baseNum] = a
			return true
		} else if aInAbc == x && xInAbcInverse == a {
			// already set but does not violate one-to-one constraint
			return true
		}

		return false
	}
	// removes chars which was set
	rmWordCharsFromSet := func(textWord, dicWord string) {
		for i := 0; i < len(textWord); i++ {
			if abc[textWord[i]-baseNum] == dicWord[i] && abcInverse[dicWord[i]-baseNum] == textWord[i] {
				abc[textWord[i]-baseNum] = 0
				abcInverse[dicWord[i]-baseNum] = 0
			}
		}
	}

	setWordChars := func(textWord, dicWord string) bool {
		for i := 0; i < len(dicWord); i++ {
			if !setChar(textWord[i], dicWord[i]) {
				return false
			}
		}
		return true
	}

	var recurMap func(twId, wId int) bool

	recurMap = func(twId, wId int) bool {
		if twId == len(textWords) {
			return true
		}

		if len(textWords[twId]) != len(words[wId]) {
			return false
		}

		if !setWordChars(textWords[twId], words[wId]) {
			rmWordCharsFromSet(textWords[twId], words[wId])
			return false
		}

		for i := 0; i < len(words); i++ {
			if recurMap(twId+1, i) {
				return true
			}
		}

		rmWordCharsFromSet(textWords[twId], words[wId])
		return false
	}

	for j := 0; j < len(words); j++ {
		if recurMap(0, j) {
			// found solution
			for i := 0; i < len(text); i++ {
				if text[i] == ' ' {
					out[i] = text[i]
				} else {
					out[i] = abc[text[i]-baseNum]
				}
			}

			return string(out)
		}
	}

	return string(out)
}

/*
3.8.2
Where’s Waldorf ?
PC/UVa IDs: 110302/10010, Popularity: B, Success rate: average Level: 2
*/
func WhereIsWaldorfFindString(grid []string, strs []string) []int {
	out := make([]int, 0, len(strs)*2)
	computeWaysLen := func(r, c int, g [][]byte) []int {
		o := make([]int, 8)
		// right
		o[0] = len(g[0]) - c
		// down right
		if o[0] > len(g)-r {
			o[1] = len(g) - r
		} else {
			o[1] = o[0]
		}
		// down
		o[2] = len(g) - r
		// down left
		if o[2] > c+1 {
			o[3] = c + 1
		} else {
			o[3] = o[2]
		}
		// left
		o[4] = c + 1
		// up left
		if o[4] > c+1 {
			o[5] = c + 1
		} else {
			o[5] = o[4]
		}
		// up
		o[6] = r + 1
		// up right
		if o[6] > len(g[0])-c {
			o[7] = len(g[0]) - c
		} else {
			o[7] = o[6]
		}
		return o
	}
	// pre process grid
	lcGrid := make([][]byte, len(grid))
	// index for inverse char search
	index := [27][][]int{}
	baseNum := byte(96)
	for i := range grid {
		lcGrid[i] = []byte(grid[i])
		for j := 0; j < len(lcGrid[i]); j++ {
			data := make([]int, 0, 10)
			ch := lcGrid[i][j]
			// ch is ascii chars
			if ch < 91 {
				ch += 32
			}

			lcGrid[i][j] = ch
			// set positions
			data = append(data, i)
			data = append(data, j)
			// compute len of all 8 ways
			ls := computeWaysLen(i, j, lcGrid)
			data = append(data, ls...)
			index[ch-baseNum] = append(index[ch-baseNum], data)
		}
	}
	lowerCaseWord := func(s string) []byte {
		o := make([]byte, len(s))
		for i := 0; i < len(s); i++ {
			if s[i] < 91 {
				o[i] = s[i] + 32
			} else {
				o[i] = s[i]
			}
		}
		return o
	}
	// check functions for each way
	checkRight := func(w []byte, r, c int, g [][]byte) bool {
		j := 0
		for i := c; i < c+len(w); i++ {
			if g[r][i] != w[j] {
				return false
			}
			j++
		}
		return true
	}
	checkDownRight := func(w []byte, r, c int, g [][]byte) bool {
		for i := 0; i < len(w); i++ {
			if g[r+i][c+i] != w[i] {
				return false
			}
		}
		return true
	}
	checkDown := func(w []byte, r, c int, g [][]byte) bool {
		j := 0
		for i := r; i < r+len(w); i++ {
			if g[i][c] != w[j] {
				return false
			}
			j++
		}
		return true
	}
	checkDownLeft := func(w []byte, r, c int, g [][]byte) bool {
		for i := 0; i < len(w); i++ {
			if g[r+i][c-i] != w[i] {
				return false
			}
		}
		return true
	}

	checkLeft := func(w []byte, r, c int, g [][]byte) bool {
		j := 0
		for i := c; i > c-len(w); i-- {
			if g[r][i] != w[j] {
				return false
			}
			j++
		}
		return true
	}
	checkUpLeft := func(w []byte, r, c int, g [][]byte) bool {
		for i := 0; i < len(w); i++ {
			if g[r-i][c-i] != w[i] {
				return false
			}
		}
		return true
	}

	checkUp := func(w []byte, r, c int, g [][]byte) bool {
		j := 0
		for i := r; i > r-len(w); i-- {
			if g[i][c] != w[j] {
				return false
			}
			j++
		}
		return true
	}

	checkUpRight := func(w []byte, r, c int, g [][]byte) bool {
		for i := 0; i < len(w); i++ {
			if g[r-i][c+i] != w[i] {
				return false
			}
		}
		return true
	}

	// match is case sensitive
	findWord := func(w []byte, index [27][][]int, grid [][]byte) []int {
		type checkFn func(w []byte, r, c int, g [][]byte) bool
		var res []int
		// load index for word's frist char
		data := index[w[0]-baseNum]
		for i := range data {
			r, c := data[i][0], data[i][1]
			checkFns := []checkFn{
				// order according to index lens computed
				checkRight, checkDownRight, checkDown, checkDownLeft,
				checkLeft, checkUpLeft, checkUp, checkUpRight,
			}
			for j, fn := range checkFns {
				// check if we have space for matching full word
				if data[i][j+2] >= len(w) && fn(w, r, c, grid) {
					res = append(res, r+1)
					res = append(res, c+1)
					// we need topmost, leftmost value
					// it will be the frist match
					break
				}
			}

		}
		return res
	}

	// main loop
	for i := 0; i < len(strs); i++ {
		w := lowerCaseWord(strs[i])
		res := findWord(w, index, lcGrid)
		// word presense expected
		out = append(out, res...)
	}

	return out
}
