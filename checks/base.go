package checks

import (
	"reflect"
	"testing"
)

func AssertEqStr(t *testing.T, s1, s2 string) {
	t.Helper()
	if s1 != s2 {
		t.Errorf("strings not equal\n%s\n%s", s1, s2)
	}
}

type Tester struct {
	t *testing.T
}

func NewTester(t *testing.T) *Tester {
	return &Tester{t}
}
func (t *Tester) AssertEq(v1, v2 interface{}) bool {
	return AssertEq(t.t, v1, v2)
}

func AssertEqErr(t *testing.T, expected, got error, msg ...string) bool {
	t.Helper()
	txt := ""
	if len(msg) > 0 {
		txt = msg[0]
	}

	if expected == got {
		return true
	}

	if expected != nil && got != nil && expected.Error() == got.Error() {
		return true
	}
	t.Errorf("%s\nErr expected\n%v\ngot\n%v\n", txt, expected, got)
	return false
}

func AssertEq(t *testing.T, expected, got interface{}, msg ...string) bool {
	t.Helper()
	txt := ""
	if len(msg) > 0 {
		txt = msg[0]
	}
	if reflect.TypeOf(expected) != reflect.TypeOf(got) {
		t.Errorf("%s\nTypes not equal Type\n%T\n%T\n", txt, expected, got)
		return false
	}
	expectedv := reflect.ValueOf(expected)
	gotv := reflect.ValueOf(got)

	if expectedv.Kind() != reflect.Slice {
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("%s\nexpected\n%v\ngot\n%v\n", txt, expected, got)
			return false
		}
		return true

	}
	expectedLen := expectedv.Len()
	if expectedLen != gotv.Len() {
		t.Errorf("%s\nslices Len not equal\n%v\n%v\n", txt, expectedLen, gotv.Len())
		return false
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("%s\nexpected\n%[2]T%[2]v\ngot\n%[3]T%[3]v\n", txt, expected, got)
		return false
	}
	return true
}
