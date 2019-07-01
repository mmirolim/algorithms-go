package graph

import (
	"fmt"
	"testing"
)

func TestSVG(t *testing.T) {
	s, e := GenSVGFromGraph(twoRhombus)
	if e != nil {
		t.Errorf("unexpected err %+v", e)
		t.FailNow()
	}
	fmt.Printf("err %+v\n", WriteToFile("example.svg", s))
}
