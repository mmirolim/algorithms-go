package tree

import (
	"errors"
	"strconv"
	"strings"
)

var ErrHeapIsFull = errors.New("heap is full")
var ErrHeapWrongSize = errors.New("wrong size")
var ErrHeapIsEmpty = errors.New("heap is empty")

type heapVal struct {
	key  int
	data interface{}
}

type Heap struct {
	size   int
	arr    []heapVal
	toSwap func(a, b int) bool
}

func newHeap(size int, toSwap func(a, b int) bool) (*Heap, error) {
	h := &Heap{}
	h.toSwap = toSwap
	if size < -1 || size == 0 {
		return nil, ErrHeapWrongSize
	}
	h.size = size
	if size > 0 {
		h.arr = make([]heapVal, 0, size)
	} else {
		h.arr = []heapVal{}
	}
	return h, nil
}

func newHeapFrom(
	heap func(int) (*Heap, error),
	keys []int,
	vals []interface{},
) (*Heap, error) {
	if vals != nil && len(keys) != len(vals) {
		return nil, errors.New("wrong number of keys or vals")
	}
	h, err := heap(len(keys))
	if err != nil {
		return nil, err
	}
	if vals == nil {
		for i := range keys {
			h.arr = append(h.arr, heapVal{keys[i], nil})
		}
	} else {
		for i := range keys {
			h.arr = append(h.arr, heapVal{keys[i], vals[i]})
		}
	}
	// heapify
	for i := len(h.arr) - 1; i >= 0; i-- {
		h.bubbleDown(i)
	}
	return h, err
}

func NewMinHeapFrom(keys []int, vals []interface{}) (*Heap, error) {
	return newHeapFrom(NewMinHeap, keys, vals)
}

func NewMaxHeapFrom(keys []int, vals []interface{}) (*Heap, error) {
	return newHeapFrom(NewMaxHeap, keys, vals)
}

func NewMinHeap(size int) (*Heap, error) {
	return newHeap(size, func(a, b int) bool {
		return a > b
	})
}

func NewMaxHeap(size int) (*Heap, error) {
	return newHeap(size, func(a, b int) bool {
		return a < b
	})
}

func (h *Heap) Insert(key int, val interface{}) error {
	if h.size != -1 && len(h.arr) >= h.size {
		return ErrHeapIsFull
	}
	h.arr = append(h.arr, heapVal{key, val})
	h.bubbleUp(len(h.arr) - 1)
	return nil
}

// extracts top element in case of MinHeap it is min
// for MaxHeap max
func (h *Heap) ExtractTop() (int, interface{}, error) {
	if len(h.arr) == 0 {
		return 0, 0, ErrHeapIsEmpty
	}
	key, val := h.arr[0].key, h.arr[0].data
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.bubbleDown(0)
	return key, val, nil
}
func (h *Heap) bubbleDown(i int) {
	c := h.firstChild(i)
	if c == -1 {
		return
	}
	id := i
	for i := 0; i < 2; i++ {
		if c+i < len(h.arr) && h.toSwap(h.arr[id].key, h.arr[c+i].key) {
			id = c + i
		}
	}
	// no need to swap
	if id == i {
		return
	}
	h.arr[i], h.arr[id] = h.arr[id], h.arr[i]
	h.bubbleDown(id)
}

func (h *Heap) firstChild(i int) int {
	id := i * 2
	if id > len(h.arr)-1 {
		return -1
	}
	return id
}

func (h *Heap) parent(i int) int {
	if i == 0 {
		return -1
	}
	return i / 2
}

func (h *Heap) bubbleUp(i int) {
	p := h.parent(i)
	if p == -1 || !h.toSwap(h.arr[p].key, h.arr[i].key) {
		return
	}
	h.arr[p], h.arr[i] = h.arr[i], h.arr[p]
	h.bubbleUp(p)
}

func (h *Heap) ToString() string {
	if len(h.arr) == 0 {
		return ""
	}

	var out strings.Builder
	for i := range h.arr {
		out.WriteString(strconv.Itoa(h.arr[i].key))
		out.WriteByte(',')
	}
	str := out.String()
	return str[:len(str)-1]
}
