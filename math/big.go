package math

// Supports only positive numbers and multiplication by int
// TODO add sign, addition, subtraction, multiplication by int and Big
type BigInt struct {
	s []uint8
}

func NewBig(n int) *BigInt {
	b := new(BigInt)
	b.s = make([]uint8, 0, 10)
	nums := parseInt(n)
	for i := 0; i < len(nums); i++ {
		b.s = append(b.s, nums[i])
	}
	return b
}

func (b *BigInt) Mul(n int) {
	nums := parseInt(n)
	adds := make([][]uint8, len(nums))
	adds[0] = b.s
	for i := 1; i < len(adds); i++ {
		adds[i] = make([]uint8, i+len(b.s))
		copy(adds[i][i:], b.s)
		scalarMul(nums[i], &adds[i])
	}
	scalarMul(nums[0], &adds[0])
	for i := 0; i < len(adds)-1; i++ {
		add(&adds[len(adds)-1], &adds[i])
	}
	b.s = adds[len(adds)-1]
}

func (b *BigInt) String() string {
	str := make([]byte, len(b.s))
	l := len(b.s)
	for i := 0; i < l; i++ {
		str[i] = b.s[l-i-1] + '0'
	}
	return string(str)
}

func parseInt(n int) []uint8 {
	out := make([]uint8, 0, 10)
	var d uint8
	for {
		d = uint8(n % 10)
		out = append(out, d)
		n = n / 10
		if n == 0 {
			break
		}
	}
	return out
}

// result is in n1
func add(n1, n2 *[]uint8) {
	s1, s2 := *n1, *n2
	if len(s1) < len(s2) {
		// grow s1
		diff := len(s2) - len(s1)
		for i := 0; i < diff; i++ {
			s1 = append(s1, 0)
		}
	}
	var carry, r, v2 uint8 = 0, 0, 0
	ls := len(s1)
	for i := 0; i < ls; i++ {
		v2 = 0
		if i < len(s2) {
			v2 = s2[i]
		}
		s1[i] = s1[i] + v2 + carry
		r = s1[i] % 10
		carry = s1[i] / 10
		s1[i] = r
	}
	if carry > 0 {
		s1 = append(s1, 1)
	}
	*n1 = s1
}

func scalarMul(c uint8, n *[]uint8) {
	if c > 9 {
		panic("c must be [0,9]")
	}
	s := *n
	var carry, r uint8 = 0, 0
	for i := 0; i < len(s); i++ {
		s[i] = s[i]*c + carry
		r = s[i] % 10
		carry = s[i] / 10
		s[i] = r
	}
	if carry > 0 {
		s = append(s, carry)
	}
	*n = s
}
