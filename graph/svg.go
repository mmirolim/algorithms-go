package graph

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/mmirolim/algos/svg"
)

const (
	defWeightFontSize = 20
	defNodeRad        = 30
	defNodeFill       = "bisque"
)

func circleIDSuffix() string {
	return "#circle"
}
func lineIDSuffix() string {
	return "#line"
}
func textIDSuffix() string {
	return "#text"
}

func lineID(v1, v2 int) string {
	return strconv.Itoa(v1) + "-" + strconv.Itoa(v2) + lineIDSuffix()
}
func lineVsID(id string) (v1, v2 int) {
	var e error
	ind := strings.IndexByte(id, '#')
	id = id[:ind]
	ind = strings.IndexByte(id, '-')
	v1, e = strconv.Atoi(id[:ind])
	if e != nil {
		panic(e)
	}
	v2, e = strconv.Atoi(id[ind+1:])
	if e != nil {
		panic(e)
	}
	return
}

func circleID(v int) string {
	return strconv.Itoa(v) + circleIDSuffix()
}

func textID(id int) string {
	return strconv.Itoa(id) + textIDSuffix()
}
func weightID(v1, v2 int) string {
	return strconv.Itoa(v1) + "-" + strconv.Itoa(v2) + "#weight"
}
func idAtoi(id string) int {
	ind := strings.IndexByte(id, '#')
	v, err := strconv.Atoi(id[:ind])
	if err != nil {
		panic("invalid id ")
	}
	return v
}

// TODO pass r, fill for nodes
func (g *Graph) ConvertToSVG(width, height, nodeSize int, dimUnitW, dimUnitH string) *svg.SVG {
	svgCon := svg.NewSVG(width, height, dimUnitW, dimUnitH)
	verticesPos := computeVerticePositions(g, 300, 300)
	nodes, labels := nodesToSvgObjs(g, verticesPos, nodeSize, defNodeFill)
	edges, weights := edgesToSvgObjs(g, nodeSize, defWeightFontSize, nodes)
	svgCon.GCircles = nodes
	svgCon.GLines = edges
	svgCon.GTexts = labels
	for k, v := range weights {
		svgCon.GTexts[k] = v
	}
	return svgCon

}

func nodesToSvgObjs(g *Graph, verticesPos [][2]int, r int, fill string) (map[string]svg.SvgCircle, map[string]svg.SvgText) {
	nodes := make(map[string]svg.SvgCircle, g.numOfVertices)
	labels := make(map[string]svg.SvgText, g.numOfVertices)
	for i := 1; i < len(g.edges); i++ {
		code := strconv.Itoa(i)
		if g.hasCodes {
			code = g.codes[i]
		}
		x, y := verticesPos[i][0], verticesPos[i][1]
		txtID, cID := textID(i), circleID(i)
		nodes[cID] = svg.NewDefSvgCircle(cID, x, y, r, fill)
		labels[txtID] = svg.NewDefSvgText(txtID, x, y, code)
	}

	return nodes, labels
}

// returns edges as lines and weights as text
func edgesToSvgObjs(
	g *Graph,
	nodeRadius, weightFontSize int,
	nodes map[string]svg.SvgCircle,
) (map[string]svg.SvgLine, map[string]svg.SvgText) {
	edges := g.Edges()
	lines := make(map[string]svg.SvgLine, g.numOfEdges)
	weights := make(map[string]svg.SvgText, g.numOfEdges)

	// compute lines
	for i := 0; i < len(edges); i++ {
		v1, v2 := edges[i][0], edges[i][1]
		n1, ok := nodes[circleID(v1)]
		if !ok {
			panic(fmt.Sprintf("edgeToLine missing node %s", circleID(v1)))
		}
		n2, ok := nodes[circleID(v2)]
		if !ok {
			panic(fmt.Sprintf("edgeToLine missing node %s", circleID(v2)))
		}
		lID := lineID(v1, v2)
		x1, y1 := n1.Coords()
		x2, y2 := n2.Coords()
		if g.isDirected {
			// compute line end on border of connected node
			r := nodeRadius
			d1 := int(Dist(x1, y1, x2, y2))
			d2 := d1 - r
			x2 = d2*(x2-x1)/d1 + x1
			y2 = d2*(y2-y1)/d1 + y1
		}
		lines[lID] = svg.NewDefSvgLine(lID, x1, y1, x2, y2, g.isDirected)
		// position weights
		x, y := MiddlePoint(x1, y1, x2, y2)
		vertex2 := g.edges[v1].find(v2)
		if vertex2.val > 0 {
			id := weightID(v1, v2)
			t := svg.NewDefSvgText(id, x, y, strconv.Itoa(vertex2.val))
			t.SetFontSize(weightFontSize)
			weights[id] = t
		}
	}

	return lines, weights
}

// compute positions for each vertices
// returns slice of [2]int{x, y} coordinates
// TODO handle directed graphs
func computeVerticePositions(g *Graph, startX, startY int) [][2]int {
	// position vertices
	pos := make([][2]int, g.numOfVertices+1)
	for i := 1; i < len(pos); i++ {
		pos[i] = [2]int{startX, startY}
	}
	verts := make([][2]int, 0, g.numOfVertices)
	for count := 0; count < 100; count++ {
		for i := 1; i < len(g.edges); i++ {
			for j := 1; j < len(g.edges); j++ {
				v1, v2 := pos[i], pos[j]
				connected := false
				if g.edges[i].find(j) != nil {
					connected = true
				}
				if !connected {
					// check maybe j connected to i in case of directed graph
					if g.edges[j].find(i) != nil {
						connected = true
					}
				}
				change, x1, y1 := pushForceDislocation(
					i, j,
					v1[0], v1[1], v2[0], v2[1],
					g.degree[i], connected)
				if change {
					verts = append(verts, [2]int{x1, y1})
				}
			}
			if len(verts) > 0 {
				x, y := verts[0][0], verts[0][1]
				origX, origY := pos[i][0], pos[i][1]
				// add all forces
				for k := 1; k < len(verts); k++ {
					x, y = AddVectorsWithSameOrigin(x, y, verts[k][0], verts[k][1], origX, origY)
				}

				pos[i][0] = x
				pos[i][1] = y
			}
			verts = verts[:0]
		}
	}

	return pos
}

// AddVectorsWithSameOrigin adds two vector with same origin
// v1 (p1, orig) v2 (p2, orig)
func AddVectorsWithSameOrigin(x1, y1, x2, y2, orx, ory int) (x, y int) {
	// move to origin
	x1, x2 = x1-orx, x2-orx
	y1, y2 = y1-ory, y2-ory
	x, y = x1+x2, y1+y2
	return x + orx, y + ory
}

// distance between points
func Dist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2.0) + math.Pow(float64(y1-y2), 2.0))
}

func MiddlePoint(x1, y1, x2, y2 int) (x, y int) {
	return x1 - (x1-x2)/2, y1 - (y1-y2)/2
}

// stretch F = kDeltaL
func attractForceByEdge(x1, y1, x2, y2 int, connected bool) float64 {
	if !connected {
		return 0.01
	}
	k := 2200.0
	defL := 100.0
	f := k * (Dist(x1, y1, x2, y2) - defL)
	return f
}

// pull force F = G/r^2
func pushForce(x1, y1, x2, y2 int) float64 {
	G := 2200.0
	return G / math.Pow(Dist(x1, y1, x2, y2), 2.0)
}

// returns dislocation by force of the first point
func pushForceDislocation(
	v1ID, v2ID int,
	x1, y1, x2, y2, c1Degree int,
	connected bool,
) (bool, int, int) {
	if v1ID == v2ID {
		return false, x1, y1
	}

	k := float64(y1-y2) / float64(x1-x2)
	f := pushForce(x1, y1, x2, y2)
	signx := 0
	signy := 0
	delta := 10
	fsign := 1
	attractForce := attractForceByEdge(x1, y1, x2, y2, connected)
	forces := f - attractForce
	// TODO refactor
	threshold := 0.03 //float64(c1Degree+1)*0.03
	if forces >= 0 {
		if forces < threshold {
			return false, x1, y1
		}
	} else {
		if forces > -threshold {
			return false, x1, y1
		}
		fsign = -1
	}
	// change location
	// find direction
	// TODO refactor
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
	x1, y1 = x1+fsign*signx*delta, y1+fsign*signy*delta
	// border force
	return true, x1 + borderForce(x1), y1 + borderForce(y1)
}

func borderForce(x int) int {
	if x > 30 {
		return 0
	}
	return 15
}
