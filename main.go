package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var args = os.Args[1:]

	if len(args) == 0 {
		log.Fatal("expected 1 argument: input file name")
	}
	var filename = args[0]

	var file, err = os.Open(filename)
	if err != nil {
		log.Fatal("could not read input file " + err.Error())
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)
	var edges = make([]Edge, 0)
	for scanner.Scan() {
		var line = scanner.Text()
		var currLineEdges, err = ReadEdgesInOrderingSequence(line)
		if err != nil {
			log.Fatal(err.Error())
		}
		edges = append(edges, currLineEdges...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err.Error())
	}

	var graph = NewGraph(edges)
	var cycle = FindCycle(graph)
	if cycle != nil {
		log.Fatalf("cannot find an ordering: cycle found\nCYCLE IS   %s\n", strings.Join(cycle, " => "))
	}

	var sorted = TopologicalSort(graph)
	fmt.Println("Resulting ordering:")
	fmt.Println(strings.Join(sorted, " => "))

}
