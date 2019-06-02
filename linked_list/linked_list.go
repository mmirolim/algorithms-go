package linked_list

import (
	"strconv"
	"strings"
)

type List struct {
	head *Node
	len  int
}

func NewList() *List {
	return &List{head: nil, len: 0}
}
func (l *List) AddOneWithStack() {
	stack := []*Node{}
	push := func(node *Node) {
		stack = append(stack, node)
	}
	pop := func() *Node {
		if len(stack) == 0 {
			return nil
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return node
	}

	current := l.head

	for current != nil {
		push(current)
		current = current.next
	}
	carry := 1
	for {
		node := pop()
		if node == nil {
			break
		}
		result := node.val + carry
		carry = result / 10
		node.val = result % 10
		if carry == 0 {
			break
		}

	}

	if carry == 1 {
		l.head = &Node{val: 1, next: l.head}
	}

}

// TODO define type for storing digits as linked list
func (l *List) AddOneRecur() {
	var addOne func(*Node) int
	addOne = func(node *Node) int {
		if node == nil {
			return 1
		}
		result := node.val + addOne(node.next)
		node.val = result % 10
		return result / 10
	}

	if addOne(l.head) == 1 {
		l.head = &Node{val: 1, next: l.head}
	}

}

func (l *List) AddOne() {
	l.Reverse()
	current := l.head
	carry := 0
	current.val++
	for current != nil {
		current.val = current.val + carry
		carry = current.val / 10
		if current.val >= 10 {
			current.val = current.val % 10
			current = current.next
		} else {
			break
		}

	}
	l.Reverse()
	if carry == 1 {
		newNode := &Node{val: 1, next: l.head}
		l.head = newNode
	}
}

func (l *List) Reverse() {
	var prev, next *Node
	current := l.head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

func (l *List) ReverseRecur() {
	var reverse func(p, c *Node) *Node
	reverse = func(prev, cur *Node) *Node {
		if cur == nil {
			return prev
		}
		next := cur.next
		cur.next = prev
		return reverse(cur, next)
	}
	l.head = reverse(nil, l.head)
}

func (l *List) Append(node *Node) {
	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

func (l *List) ToString() string {
	var str strings.Builder
	str.WriteString("head->")
	if l.head == nil {
		str.WriteString("nil")
		return str.String()
	}
	current := l.head
	for {
		str.WriteString(strconv.Itoa(current.val))
		str.WriteString("->")
		if current.next != nil {
			current = current.next
		} else {
			break
		}
	}
	str.WriteString("nil")
	return str.String()
}

// problem ADM 3.27
func (l *List) FindAndFixLoop() *Node {
	if l == nil || l.head == nil {
		return nil
	}
	// find loop
	slow, fast := l.head, l.head
	for {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			break
		}
		if fast == nil {
			return nil
		}
	}

	var node *Node
	slow = l.head
	for {
		slow = slow.next
		if fast.next == slow {
			break
		}
		fast = fast.next
	}
	node = fast.next
	fast.next = nil
	return node
}
func (l *List) FindAndFixLoopWithMeter() *Node {
	if l == nil || l.head == nil {
		return nil
	}
	slow, fast := l.head, l.head
	for {
		slow = slow.next
		fast = fast.next.next
		if fast == nil {
			return nil
		}
		if fast == slow {
			break
		}
	}

	// loop size
	k := 0
	for {
		slow = slow.next
		k++
		if slow == fast {
			break
		}
	}

	slow, fast = l.head, l.head
	for k > 0 {
		fast = fast.next
		k--
	}

	for {
		slow = slow.next
		if fast.next == slow {
			break
		}
		fast = fast.next
	}

	node := fast.next
	fast.next = nil
	return node
}

type Node struct {
	val  int
	next *Node
}

func NewNode(val int) *Node {
	return &Node{val: val, next: nil}
}

func FindLLMiddle(ll *List) *Node {
	if ll.head == nil {
		return nil
	} else if ll.head.next == nil {
		return ll.head
	}

	n1, n2 := ll.head, ll.head
	for {
		n1 = n1.next
		n2 = n2.next.next
		if n2.next == nil {
			break
		}
	}

	return n1
}
