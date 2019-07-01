package graph

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	defFontSize   = 14
	defTextFill   = "black"
	defTextAnchor = "middle"
	defStroke     = "black"
)

var ErrSVGNodeNotFound = errors.New("err svg node not found for an edge")

func circleIDSuffix() string {
	return "#circle"
}
func lineIDSuffix() string {
	return "#line"
}
func textIDSuffix() string {
	return "#text"
}

type SvgCircle struct {
	id      string
	x, y, r int
	fill    string
	stroke  string
}

func NewDefSvgCircle(id string, x, y, r int, fill string) SvgCircle {
	return SvgCircle{
		id:     id + circleIDSuffix(),
		x:      x,
		y:      y,
		r:      r,
		fill:   fill,
		stroke: defStroke,
	}
}
func (c *SvgCircle) serialize() string {
	return fmt.Sprintf("<circle cx=\"%d\" cy=\"%d\" r=\"%d\" fill=\"%s\" stroke=\"%s\" stroke-width=\"1\"></circle>", c.x, c.y, c.r, c.fill, c.stroke)
}

type GroupCircles struct {
	class   string
	circles []SvgCircle
}
type SvgLine struct {
	id             string
	x1, y1, x2, y2 int
	stroke         string
}

func (l *SvgLine) serialize() string {
	return fmt.Sprintf("<line x1=\"%d\" y1=\"%d\" x2=\"%d\" y2=\"%d\" stroke=\"%s\"></line>",
		l.x1, l.y1, l.x2, l.y2, l.stroke)
}
func NewDefSvgLine(id string, x1, y1, x2, y2 int) SvgLine {
	return SvgLine{
		id:     id + lineIDSuffix(),
		x1:     x1,
		y1:     y1,
		x2:     x2,
		y2:     y2,
		stroke: defStroke,
	}
}

type GroupLines struct {
	class string
	lines []SvgLine
}
type SvgText struct {
	id         string
	x, y       int
	fontSize   int
	fill       string
	textAnchor string
	text       string
}

func NewSvgText(id string, x, y, fontSize int, text, fill, textAnchor string) SvgText {
	return SvgText{
		id:         id + textIDSuffix(),
		x:          x,
		y:          y,
		fontSize:   fontSize,
		fill:       fill,
		textAnchor: textAnchor,
		text:       text,
	}
}

func (t *SvgText) serialize() string {
	return fmt.Sprintf("<text x=\"%d\" y=\"%d\" font-size=\"%d\" fill=\"%s\" text-anchor=\"%s\">%s</text>",
		t.x, t.y, t.fontSize, t.fill, t.textAnchor, t.text)
}

func NewDefSvgText(id string, x, y int, text string) SvgText {
	return SvgText{
		id:         id + textIDSuffix(),
		x:          x,
		y:          y,
		text:       text,
		fontSize:   defFontSize,
		fill:       defTextFill,
		textAnchor: defTextAnchor,
	}
}

type GroupTexts struct {
	texts []SvgText
}
type SVG struct {
	g                              *Graph
	width, height                  int
	dimensionUnitW, dimensionUnitH string
	defFontSize                    int
	defTextFill                    string
	defTextAnchor                  string
	defStroke                      string
	GLines                         map[string]SvgLine
	GCircles                       map[string]SvgCircle
	GTexts                         map[string]SvgText
}

func NewSVG(g *Graph, width, height int, dimensionUnitW, dimensionUnitH string) (*SVG, error) {
	if width == 0 {
		width = 100
	}
	if height == 0 {
		height = 100
	}
	// TODO check valid units
	if dimensionUnitW == "" {
		dimensionUnitW = "%"
	}
	if dimensionUnitH == "" {
		dimensionUnitH = "%"
	}
	svgCon := &SVG{
		g:              g,
		width:          width,
		height:         height,
		dimensionUnitW: dimensionUnitW,
		dimensionUnitH: dimensionUnitH,
		defFontSize:    14,
		defTextFill:    "#999",
		defTextAnchor:  "middle",
		defStroke:      "black",
		GLines:         map[string]SvgLine{},
		GCircles:       map[string]SvgCircle{},
		GTexts:         map[string]SvgText{},
	}
	err := svgCon.readGraph()
	if err != nil {
		return nil, err
	}
	return svgCon, nil
}

func (s *SVG) readGraph() error {
	//	maxWidth := len(s.g.edges) * 100
	//	maxHeight := len(s.g.edges) * 100
	for i := 1; i < len(s.g.edges); i++ {
		code := strconv.Itoa(i)
		if s.g.hasCodes {
			code = s.g.codes[i]
		}
		c, txt := s.vertexToCircleWithText(
			i, code, 300, 300, 30, "bisque")
		s.GCircles[c.id] = c
		s.GTexts[txt.id] = txt
	}
	// position nodes
	updateLoc := true
	count := 0
	edges := s.g.Edges()
	for updateLoc {
		count++
		if count > 100 {
			break
		}
		updateLoc = false
		vecs := make([][2]int, 0, len(s.GCircles))
		id := 0
		for k1 := range s.GCircles {
			change, x1, y1 := false, 0, 0
			for k2 := range s.GCircles {
				c1, c2 := s.GCircles[k1], s.GCircles[k2]
				id = idAtoi(c1.id)
				connected := false
				if s.g.edges[id].find(idAtoi(c2.id)) != nil {
					connected = true
				}
				change, x1, y1 = pushForceDislocation(&c1, &c2, s.g.degree[id], connected)
				if change {
					vecs = append(vecs, [2]int{x1, y1})
					updateLoc = true
				}
			}
			if change {
				c1 := s.GCircles[k1]
				x, y := vecs[0][0], vecs[0][1]
				for i := 1; i < len(vecs); i++ {
					x, y = addVector(x, y, vecs[i][0], vecs[i][1], c1.x, c1.y)
				}

				c1.x, c1.y = x, y
				s.GCircles[k1] = c1

			}
			vecs = vecs[:0]
		}
	}
	// compute lines
	for i := 0; i < len(edges); i++ {
		l, err := s.edgeToLine(edges[i][0], edges[i][1])
		if err != nil {
			return err
		}
		s.GLines[l.id] = l
	}
	// position text
	for k, c := range s.GCircles {
		tid := strconv.Itoa(idAtoi(k)) + textIDSuffix()
		t := s.GTexts[tid]
		t.x, t.y = c.x, c.y
		s.GTexts[tid] = t
	}
	return nil
}
func idAtoi(id string) int {
	ind := strings.IndexByte(id, '#')
	v, err := strconv.Atoi(id[:ind])
	if err != nil {
		panic("invalid id ")
	}
	return v
}

func (s *SVG) SetDefFontSize(size int) *SVG {
	s.defFontSize = size
	return s
}

func (s *SVG) SetDefTextFill(fill string) *SVG {
	s.defTextFill = fill
	return s
}
func (s *SVG) SetDefTextAnchor(pos string) *SVG {
	s.defTextAnchor = pos
	return s
}
func (s *SVG) SetDefStroke(stroke string) *SVG {
	s.defStroke = stroke
	return s
}

func (s *SVG) vertexToCircleWithText(
	id int,
	label string,
	x, y, r int,
	fill string,
) (SvgCircle, SvgText) {
	idStr := strconv.Itoa(id)
	c := NewDefSvgCircle(idStr, x, y, r, fill)
	if label == "" {
		label = strconv.Itoa(id)
	}
	txt := NewDefSvgText(idStr, x, y, label)
	return c, txt
}

// TODO add weight support
func (s *SVG) edgeToLine(v1, v2 int) (SvgLine, error) {
	v1Str := strconv.Itoa(v1)
	v2Str := strconv.Itoa(v2)
	idStr := v1Str + v2Str
	n1, ok := s.GCircles[v1Str+circleIDSuffix()]
	if !ok {
		return SvgLine{}, ErrSVGNodeNotFound
	}
	n2, ok := s.GCircles[v2Str+circleIDSuffix()]
	if !ok {
		return SvgLine{}, ErrSVGNodeNotFound
	}
	l := NewDefSvgLine(idStr, n1.x, n1.y, n2.x, n2.y)
	return l, nil
}

func addVector(x1, y1, x2, y2, orx, ory int) (x, y int) {
	// move to origin
	x1, x2 = x1-orx, x2-orx
	y1, y2 = y1-ory, y2-ory
	x, y = x1+x2, y1+y2
	return x + orx, y + ory
}

// F = kDeltaL
func attractForceByEdge(c1, c2 *SvgCircle, connected bool) float64 {
	if !connected {
		return 0.0
	}
	x1, y1, x2, y2 := c1.x, c1.y, c2.x, c2.y
	k := 2200.0
	defL := 100.0
	f := k * (math.Sqrt(math.Pow(float64(x1-x2), 2.0)+math.Pow(float64(y1-y2), 2.0)) - defL)
	return f
}

// F = G/r^2
func pushForce(c1, c2 *SvgCircle, c1Degree int) float64 {
	x1, y1, x2, y2 := c1.x, c1.y, c2.x, c2.y
	G := 2200.0
	f := G / (math.Pow(float64(x1-x2), 2.0) + math.Pow(float64(y1-y2), 2.0))
	return f
}

func pushForceDislocation(c1, c2 *SvgCircle, c1Degree int, connected bool) (bool, int, int) {
	if c1.id == c2.id {
		return false, c1.x, c1.y
	}
	x1, y1, x2, y2 := c1.x, c1.y, c2.x, c2.y
	k := float64(y1-y2) / float64(x1-x2)
	f := pushForce(c1, c2, c1Degree)
	// TODO move to separate function
	signx := 0
	signy := 0
	delta := 10
	fsign := 1
	forces := f - attractForceByEdge(c1, c2, connected)
	if forces >= 0 {
		if forces < float64(c1Degree+1)*0.03 {
			return false, x1, y1
		}
	} else {
		fsign = -1
	}
	// change location
	if x1 < x2 {
		signx = -1
		if k > 0 {
			signy = -1
		} else {
			signy = 1
		}
	} else if x1 == x2 {
		if y1 > y2 {
			signy = 1
		} else if y1 < y2 {
			signy = -1
		} else {
			// special case nodes in same pos
			// randomly assing dir
			if rand.Float64() > 0.5 {
				signx = 1
			} else {
				signx = -1
			}
			if rand.Float64() > 0.5 {
				signy = 1
			} else {
				signy = -1
			}
		}

	} else {
		signx = 1
		if k > 0 {
			signy = 1
		} else {
			signy = -1
		}
	}
	return true, x1 + fsign*signx*delta, y1 + fsign*signy*delta

}

func (s *SVG) serialize() string {
	var str strings.Builder
	header := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" ?>
<svg xmlns="http://www.w3.org/2000/svg" width="%d%s" height="%d%s">`, s.width, s.dimensionUnitW,
		s.height, s.dimensionUnitH)
	footer := `</svg>`
	if s.g == nil {
		return header + footer
	}
	str.WriteString(header)
	str.WriteByte('\n')
	// next element cover prev
	// first lines
	for _, v := range s.GLines {
		str.WriteString(v.serialize())
		str.WriteByte('\n')
	}
	// second nodes
	for _, v := range s.GCircles {
		str.WriteString(v.serialize())
		str.WriteByte('\n')
	}
	// third text
	for _, v := range s.GTexts {
		str.WriteString(v.serialize())
		str.WriteByte('\n')
	}
	str.WriteString(footer)
	return str.String()
}

var graphData = `
5 4 directed weighted hascodes
1 Bill Clinton
2 Hillary Clinton
3 John McCain
4 George Bush
5 Saddam Hussein
1 2 10
1 3 2
2 3 5
3 4 4
5
`
var twoRhombus = `7 8
1 3
1 2
2 4
3 4
4 5
4 7
5 6
6 7
`

func GenSVGFromGraph(data string) (*SVG, error) {
	g, e := NewGraphFrom(data)
	if e != nil {
		return nil, e
	}
	svgCon, e := NewSVG(g, 100, 100, "%", "%")
	if e != nil {
		return nil, e
	}
	return svgCon, nil
}

func WriteToFile(fname string, svgCon *SVG) error {
	f, e := os.Create(fname)
	if e != nil {
		return e
	}
	defer f.Close()
	_, e = f.WriteString(svgCon.serialize())
	return e
}
