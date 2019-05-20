package linked_list

import (
	"strconv"
	"strings"
)

var list1 = &Node{val: 5, next: &Node{val: 6, next: &Node{val: 3, next: nil}}}
var list2 = &Node{val: 8, next: &Node{val: 4, next: &Node{val: 2, next: nil}}}

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

type Node struct {
	val  int
	next *Node
}

func NewNode(val int) *Node {
	return &Node{val: val, next: nil}
}
