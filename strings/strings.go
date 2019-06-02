package strings

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
