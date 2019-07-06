package math

// Supports only positive numbers and multiplication by int
// TODO add sign, addition, subtraction, multiplication by int and Big
type BigInt struct {
	s     []uint8
	nums  []uint8
	temp  []uint8
	final []uint8
}

func NewBig(n int) *BigInt {
	b := new(BigInt)
	b.s = make([]uint8, 0, 10)
	b.nums = []uint8{}
	b.temp = make([]uint8, 10)
	b.final = make([]uint8, 10)
	parseInt(n, &b.nums)
	for i := 0; i < len(b.nums); i++ {
		b.s = append(b.s, b.nums[i])
	}
	b.nums = b.nums[:0]
	return b
}

func (b *BigInt) Mul(n int) {
	parseInt(n, &b.nums)
	maxLen := len(b.s) + len(b.nums)
	if maxLen > len(b.temp) {
		// extend
		b.temp = append(b.temp,
			make([]uint8, maxLen-len(b.temp))...)
	} else {
		b.temp = b.temp[:maxLen]
	}
	for i := 0; i < len(b.final); i++ {
		b.final[i] = 0
	}

	if maxLen > len(b.final) {
		// extend
		b.final = append(b.final,
			make([]uint8, maxLen-len(b.final))...)
	} else {
		b.final = b.final[:maxLen]
	}

	for i := 0; i < len(b.nums); i++ {
		for j := 0; j < i; j++ {
			b.temp[j] = 0
		}
		tmp := b.temp[i:]
		scalarMul(&tmp, b.nums[i], b.s)
		add(&b.final, &b.temp)
	}

	last := len(b.final)
	var prev uint8
	for i := len(b.final) - 1; i > 0; i-- {
		if prev == 0 && b.final[i] == 0 {
			last = i
		} else {
			break
		}
		prev = b.final[i]
	}
	b.nums = b.nums[:0]
	b.s, b.final = b.final[:last], b.s
}

func (b *BigInt) String() string {
	str := make([]byte, len(b.s))
	l := len(b.s)
	for i := 0; i < l; i++ {
		str[i] = b.s[l-i-1] + '0'
	}
	return string(str)
}

func parseInt(n int, nums *[]uint8) {
	ds := *nums
	var d uint8
	for {
		d = uint8(n % 10)
		ds = append(ds, d)
		n = n / 10
		if n == 0 {
			break
		}
	}
	*nums = ds
}

func add(n1, n2 *[]uint8) {
	s1, s2 := *n1, *n2
	if len(s1) < len(s2) {
		// grow s1
		diff := len(s2) - len(s1)
		for i := 0; i < diff; i++ {
			s1 = append(s1, 0)
		}
	}
	var carry uint8
	ls := len(s2)
	for i := 0; i < ls; i++ {
		s1[i] = s1[i] + s2[i] + carry
		carry = s1[i] / 10
		s1[i] = s1[i] % 10

	}
	for carry > 0 {
		s1[ls] = s1[ls] + carry
		carry = s1[ls] / 10
		s1[ls] = s1[ls] % 10
		ls++
	}
	*n1 = s1
}

func scalarMul(dist *[]uint8, c uint8, s []uint8) {
	if c > 9 {
		panic("c must be [0,9]")
	}
	d := *dist
	var carry uint8
	ls := len(s)
	for i := 0; i < ls; i++ {
		d[i] = s[i]*c + carry
		carry = d[i] / 10
		d[i] = d[i] % 10
	}

	if carry > 0 {
		d[ls] = d[ls] + carry
	}
	*dist = d
}
