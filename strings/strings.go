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
		abc := [25]int{}
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
