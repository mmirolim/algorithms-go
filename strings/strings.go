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
