package graph

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var (
	ErrWrongNumOfVertices = errors.New("err wrong number of vertices")
	ErrWeightMissing      = errors.New("err weight missing")
	ErrWrongNumOfEdges    = errors.New("err wrong number of edges")
)

type Graph struct {
	isDirected, isWeighted, hasCodes bool
	numOfEdges                       int
	numOfVertices                    int

	// Adjacency list to store edges
	edges []*list
	// degree of each vertex
	degree []int
	codes  map[int]string
}

// NewGraphFrom creates graph from string
// space separated vertex, edge directed keyword, weighted keyword
// next line space separated code and name of vertex
// next line space separated vertices of edge and weight it exists
// edges code should start from 1 and sequential order
func NewGraphFrom(txt string) (*Graph, error) {
	g := new(Graph)
	txt = strings.Trim(txt, " \t\r\n")
	r := strings.NewReader(txt)

	// readline
	line, err := readLine(r)
	if err != nil {
		return nil, err
	}
	err = g.graphCfgFromString(line)
	if err != nil {
		return nil, err
	}
	// init storage
	g.edges = make([]*list, g.numOfVertices+1)
	g.degree = make([]int, g.numOfVertices+1)
	// read codes
	if g.hasCodes {
		for i := 0; i < g.numOfVertices; i++ {
			line, err = readLine(r)
			if err != nil {
				return nil, err
			}
			g.codeFromString(line)
		}
	}

	// store all seen vertices
	seenVertices := map[int]bool{}
	// read edges
	for err == nil {
		line, err = readLine(r)
		if err != nil {
			if err != io.EOF {
				return nil, err
			} else if line == "" {
				break
			}
		}
		err = g.edgeFromString(line, seenVertices)
		if err != nil {
			return nil, err
		}
	}
	err = checkNumOfVertices(g.numOfVertices, seenVertices)
	if err != nil {
		return nil, err
	}
	err = g.checkNumOfEdges()
	if err != nil {
		return nil, err
	}
	return g, nil
}
func (g *Graph) checkNumOfEdges() error {
	for i := range g.edges {
		if g.edges[i] != nil {
			g.degree[i] = g.edges[i].len
		}
	}
	edges := g.Edges()
	if len(edges) != g.numOfEdges {
		return ErrWrongNumOfEdges
	}
	return nil
}
func checkNumOfVertices(numOfVertices int, seenVertices map[int]bool) error {
	if len(seenVertices) != numOfVertices {
		return ErrWrongNumOfVertices
	}
	min, max := 1<<30, -(1 << 30)
	for k := range seenVertices {
		if k < min {
			min = k
		}
		if k > max {
			max = k
		}
	}
	if min != 1 || max != numOfVertices {
		return ErrWrongNumOfVertices
	}
	return nil
}

func (g *Graph) graphCfgFromString(line string) error {
	var err error
	parts := strings.SplitN(string(line), " ", 3)
	g.numOfVertices, err = strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	g.numOfEdges, err = strconv.Atoi(parts[1])
	if err != nil {
		return err
	}

	if len(parts) == 3 {
		if strings.Contains(parts[2], "directed") {
			g.isDirected = true
		}
		if strings.Contains(parts[2], "weighted") {
			g.isWeighted = true
		}
		if strings.Contains(parts[2], "hascodes") {
			g.hasCodes = true
			g.codes = make(map[int]string, g.numOfVertices)
		}
	}

	return nil
}

func (g *Graph) codeFromString(line string) error {
	spaceIndex := strings.IndexByte(line, ' ')
	code, err := strconv.Atoi(line[:spaceIndex])
	if err != nil {
		return err
	}
	g.codes[code] = line[spaceIndex:]
	return nil
}

func (g *Graph) edgeFromString(line string, seenVertices map[int]bool) error {
	var v1, v2, w int
	var err error
	if strings.Count(line, " ") == 0 {
		// single degree vertex
		fmt.Sscanf(line, "%d\n", &v1)
		seenVertices[v1] = true
		return nil
	}
	if g.isWeighted {
		_, err = fmt.Sscanf(line, "%d %d %d\n", &v1, &v2, &w)
		// if got unexpected EOF must be weight missing
		if err == io.EOF {
			return ErrWeightMissing
		}

	} else {
		_, err = fmt.Sscanf(line, "%d %d\n", &v1, &v2)
		w = 0
	}
	seenVertices[v1], seenVertices[v2] = true, true
	if err != nil {
		if err != io.EOF {
			return err
		}
	}
	g.insertEdge(v1, v2, w, g.isDirected)
	return nil
}

func (g *Graph) insertEdge(v1, v2, w int, directed bool) {
	if g.edges[v1] == nil {
		g.edges[v1] = newList(newNode(v2, w))
	} else {
		g.edges[v1].append(newNode(v2, w))
	}
	if !directed {
		g.insertEdge(v2, v1, w, true)
	}
}

func (g *Graph) IsDirected() bool {
	return g.isDirected
}
func (g *Graph) IsWeighted() bool {
	return g.isWeighted
}

func (g *Graph) Vertices() []int {
	var vertices []int
	for i := 1; i < len(g.edges); i++ {
		vertices = append(vertices, i)
	}
	return vertices
}

func (g *Graph) NumOfEdges() int {
	return g.numOfEdges
}

func (g *Graph) NumOfVertices() int {
	return g.numOfVertices
}

func (g *Graph) Edges() [][2]int {
	var edges [][2]int
	v2v1Map := map[int]map[int]bool{}
	for i := 1; i < len(g.edges); i++ {
		listOfVertices := g.edges[i].ids()
		for _, j := range listOfVertices {
			if g.isSameEdge(i, j, v2v1Map) {
				continue
			}
			edges = append(edges, [2]int{i, j})
		}
	}
	return edges
}

func (g *Graph) EdgesWithWeights() [][3]int {
	var edges [][3]int
	v2v1Map := map[int]map[int]bool{}
	for i := 1; i < len(g.edges); i++ {
		listOfVertices := g.edges[i].ids()
		listOfWeights := g.edges[i].weights()
		for j := range listOfVertices {
			if g.isSameEdge(i, listOfVertices[j], v2v1Map) {
				continue
			}
			edges = append(edges, [3]int{i, listOfVertices[j], listOfWeights[j]})
		}
	}
	return edges
}
func (g *Graph) isSameEdge(v1, v2 int, m map[int]map[int]bool) bool {
	if !g.isDirected {
		if m[v2][v1] {
			return true
		}
		if m[v1] == nil {
			m[v1] = map[int]bool{}
		}
		m[v1][v2] = true
	}
	return false
}

func (g *Graph) ToString() string {
	var str strings.Builder
	str.WriteString(strconv.Itoa(g.numOfVertices))
	str.WriteString(" ")
	str.WriteString(strconv.Itoa(g.numOfEdges))
	if g.isDirected {
		str.WriteString(" directed")
	}
	if g.isWeighted {
		str.WriteString(" weighted")
	}
	if len(g.codes) > 0 {
		str.WriteString(" hascodes")
	}
	str.WriteString("\n")
	// add codes
	codes := make([]string, len(g.codes)+1)
	if g.hasCodes {
		for k := range g.codes {
			codes[k] = g.codes[k]
		}
		for i := 1; i < len(codes); i++ {
			str.WriteString(strconv.Itoa(i))
			str.WriteByte(' ')
			str.WriteString(codes[i])
			str.WriteByte('\n')
		}
	}
	// add edges
	v2v1Map := map[int]map[int]bool{}
	for i := 1; i < len(g.edges); i++ {
		ids := g.edges[i].ids()
		if len(ids) == 0 {
			str.WriteString(strconv.Itoa(i))
			str.WriteByte('\n')
		} else {
			weights := g.edges[i].weights()
			for j := 0; j < len(ids); j++ {
				if g.isSameEdge(i, ids[j], v2v1Map) {
					continue
				}
				str.WriteString(strconv.Itoa(i))
				str.WriteByte(' ')
				str.WriteString(strconv.Itoa(ids[j]))
				if g.isWeighted {
					str.WriteByte(' ')
					str.WriteString(strconv.Itoa(weights[j]))
				}
				str.WriteByte('\n')
			}
		}
	}

	return str.String()
}

type list struct {
	len  int
	head *node
	last *node
}

type node struct {
	id   int
	val  int
	next *node
}

func newNode(id, val int) *node {
	return &node{id, val, nil}
}

func newList(n *node) *list {
	if n != nil {
		return &list{1, n, n}
	}
	return &list{0, nil, nil}
}

// const time
func (l *list) prepend(n *node) {
	if n == nil {
		return
	}
	l.len++
	n.next = l.head
	l.head = n
}

// const time
func (l *list) append(n *node) {
	if n == nil {
		return
	}
	l.len++
	if l.head == nil {
		l.head = n
		l.last = n
		return
	}
	l.last.next = n
	l.last = n
}

func (l *list) find(id int) *node {
	if l.head == nil {
		return nil
	}
	current := l.head
	for current != nil {
		if current.id == id {
			return current
		}
		current = current.next
	}
	return nil
}

func (l *list) walk(f func(*node)) {
	if l == nil {
		return
	}
	current := l.head
	for current != nil {
		f(current)
		current = current.next
	}
}

func (l *list) ids() []int {
	var ids []int
	l.walk(func(n *node) {
		ids = append(ids, n.id)
	})
	return ids
}

func (l *list) weights() []int {
	var ws []int
	l.walk(func(n *node) {
		ws = append(ws, n.val)
	})
	return ws

}

func (l *list) ToString() string {
	var str strings.Builder
	if l.head == nil {
		return ""
	}
	str.WriteString("-")
	current := l.head
	for {
		str.WriteString(strconv.Itoa(current.val))
		str.WriteString("- ")
		str.WriteString(strconv.Itoa(current.id))
		str.WriteString("- ")
		if current.next != nil {
			current = current.next
		} else {
			break
		}
	}
	return str.String()
}

func readLine(r io.Reader) (string, error) {
	// readline
	var line []byte
	var b = make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err != nil {
			return string(line), err
		}
		if b[0] == '\n' {
			break
		}
		line = append(line, b[0])
	}
	return string(line), nil
}
