package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var arrowMatcher regexp.Regexp

func init() {
	arrowMatcher = *regexp.MustCompile("(=*)>(>*)")
}

func ReadOrderingSequence(str string) ([]Vertex, error) {
	var sp = strings.Split(str, " ")
	var vertices = make([]Vertex, 0, len(sp)/2+1)
	var expectVertex = true
	for _, word := range sp {
		switch {
		case expectVertex && isVertexIdentifier(word):
			vertices = append(vertices, word)
			expectVertex = false
		case expectVertex:
			return nil, errors.New(fmt.Sprintf("expected a name, found %s\n", word))
		case arrowMatcher.Match([]byte(word)):
			expectVertex = true
		case len(word) > 0:
			return nil, errors.New(fmt.Sprintf("expected an arrow (=>, ==>, ===>>, etc.), found %s\n", word))
		}
	}
	return vertices, nil
}

func ReadEdgesInOrderingSequence(str string) ([]Edge, error) {
	var verticesSeq, err = ReadOrderingSequence(str)
	if err != nil {
		return nil, err
	}
	var edges = make([]Edge, 0, len(verticesSeq)-1)
	var prevVertex Vertex = ""
	for _, currVertex := range verticesSeq {
		if len(prevVertex) > 0 {
			edges = append(edges, NewEdge(prevVertex, currVertex))
		}
		prevVertex = currVertex
	}
	return edges, nil
}

func isVertexIdentifier(str string) bool {
	for _, char := range str {
		if !(unicode.IsLetter(char) || unicode.IsDigit(char) || char == '_' || char == '-') {
			return false
		}
	}
	return true
}
