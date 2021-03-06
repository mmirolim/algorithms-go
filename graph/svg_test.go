package graph

import (
	"io/ioutil"
	"testing"
)

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
