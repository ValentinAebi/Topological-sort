package main

import (
	"reflect"
	"testing"
)

func TestReadOrderingSequence(t *testing.T) {
	input := "A ==> B ==> Foo => 2 ==>>> Bar12 ====> 215"
	act, err := ReadOrderingSequence(input)
	if err != nil {
		t.Fatal("returned error")
	}
	exp := []Vertex{"A", "B", "Foo", "2", "Bar12", "215"}
	if !reflect.DeepEqual(exp, act) {
		t.Fatalf("\nExpected: %s\n  Actual: %s\n", exp, act)
	}
}
