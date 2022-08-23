package main

import (
	"reflect"
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	A, B, C, D, E, F, G := "A", "B", "C", "D", "E", "F", "G"
	graph := NewGraph([]Edge{
		NewEdge(A, E),
		NewEdge(A, C),
		NewEdge(E, C),
		NewEdge(C, B),
		NewEdge(B, D),
		NewEdge(F, D),
		NewEdge(F, A),
		NewEdge(G, F),
		NewEdge(G, E),
	})
	exp := []string{G, F, A, E, C, B, D}
	act := TopologicalSort(graph)
	if !reflect.DeepEqual(exp, act) {
		t.Fatalf("\nExpected: %s\n  Actual: %s\n", exp, act)
	}
}
