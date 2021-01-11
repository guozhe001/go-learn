package main

import (
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJoinWhtiCommas(t *testing.T) {
	testDatas := []testData{{[]string{"apple", "orange"}, "apple and orange"},
		{[]string{"apple"}, "apple"},
		{[]string{"apple", "orange", "banana"}, "apple, orange, and banana"},
		{[]string{}, ""}}
	for _, data := range testDatas {
		invokeAndAssert(t, data.list, data.want)
	}
}

func invokeAndAssert(t *testing.T, list []string, want string) {
	got := JoinWhtiCommas(list)
	if want != got {
		t.Errorf("JoinWhtiCommas(%#v) = \"%s\" want \"%s\"", list, got, want)
	}
}
