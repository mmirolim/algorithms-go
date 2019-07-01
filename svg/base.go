package svg

import (
	"errors"
	"fmt"
	"strings"
)

const (
	defFontSize   = 14
	defTextFill   = "black"
	defTextAnchor = "middle"
	defStroke     = "black"
)

var ErrSVGNodeNotFound = errors.New("err svg node not found for an edge")

type SvgCircle struct {
	id      string
	x, y, r int
	fill    string
	stroke  string
}

func (c *SvgCircle) Coords() (x, y int) {
	return c.x, c.y
}
func (c *SvgCircle) Radius() int {
	return c.r
}

func NewDefSvgCircle(id string, x, y, r int, fill string) SvgCircle {
	return SvgCircle{
		id:     id,
		x:      x,
		y:      y,
		r:      r,
		fill:   fill,
		stroke: defStroke,
	}
}
func (c *SvgCircle) serialize() string {
	return fmt.Sprintf("<circle id=\"%s\" cx=\"%d\" cy=\"%d\" r=\"%d\" fill=\"%s\" stroke=\"%s\" stroke-width=\"1\"></circle>", c.id, c.x, c.y, c.r, c.fill, c.stroke)
}

type GroupCircles struct {
	class   string
	circles []SvgCircle
}
type SvgLine struct {
	id             string
	x1, y1, x2, y2 int
	stroke         string
	isDirected     bool
}

func (l *SvgLine) Coords() (x1, y1, x2, y2 int) {
	return l.x1, l.y1, l.x2, l.y2
}

func (l *SvgLine) serialize() string {
	marker := `marker-end="url(#arrow)"`
	if !l.isDirected {
		marker = ""
	}
	return fmt.Sprintf(`<line id="%s" x1="%d" y1="%d" x2="%d" y2="%d" stroke="%s" %s></line>`,
		l.id, l.x1, l.y1, l.x2, l.y2, l.stroke, marker)
}

// directed from v1 to v2
func NewDefSvgLine(id string, x1, y1, x2, y2 int, directed bool) SvgLine {
	return SvgLine{
		id:         id,
		x1:         x1,
		y1:         y1,
		x2:         x2,
		y2:         y2,
		isDirected: directed,
		stroke:     defStroke,
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
		id:         id,
		x:          x,
		y:          y,
		fontSize:   fontSize,
		fill:       fill,
		textAnchor: textAnchor,
		text:       text,
	}
}
func (t *SvgText) SetFontSize(size int) {
	t.fontSize = size
}
func (t *SvgText) serialize() string {
	return fmt.Sprintf("<text id=\"%s\" x=\"%d\" y=\"%d\" font-size=\"%d\" fill=\"%s\" text-anchor=\"%s\">%s</text>",
		t.id, t.x, t.y, t.fontSize, t.fill, t.textAnchor, t.text)
}

func NewDefSvgText(id string, x, y int, text string) SvgText {
	return SvgText{
		id:         id,
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

func NewSVG(width, height int, dimensionUnitW, dimensionUnitH string) *SVG {
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
	return svgCon
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

func (s *SVG) Serialize() string {
	var str strings.Builder
	header := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" ?>
<svg xmlns="http://www.w3.org/2000/svg" width="%d%s" height="%d%s">
<defs>
    <!-- arrowhead marker definition -->
    <marker id="arrow" viewBox="0 0 10 10" refX="5" refY="5"
        markerWidth="10" markerHeight="10"
        orient="auto-start-reverse">
      <path d="M 0 0 L 10 5 L 0 10 z" />
    </marker>
</defs>`, s.width, s.dimensionUnitW,
		s.height, s.dimensionUnitH)
	footer := `</svg>`
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
