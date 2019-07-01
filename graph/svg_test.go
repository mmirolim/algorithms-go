package graph

import (
	"io/ioutil"
	"testing"
)

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

func TestConvertToSVG(t *testing.T) {
	g, e := NewGraphFrom(graphDataDirectedWeightedWithCodes)
	if e != nil {
		t.Errorf("NewGraphFrom unexpected err %+v", e)
		t.FailNow()
	}
	s := g.ConvertToSVG(100, 100, 30, "%", "%")
	e = ioutil.WriteFile("example.svg", []byte(s.Serialize()), 0666)
	if e != nil {
		t.Errorf("NewGraphFrom unexpected err %+v", e)
		t.FailNow()
	}
}
