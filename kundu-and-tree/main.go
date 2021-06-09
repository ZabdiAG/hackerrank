package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Color color edge
type Color uint8

const (
	// ColorBlue blue
	ColorBlue Color = iota + 1
	// ColorRed red
	ColorRed
)

// Edge ...
type Edge struct {
	Color  Color
	Vertex [2]*Vertex
}

// Vertex ...
type Vertex struct {
	ID    int64
	Edges []*Edge
}

// MapVertex ...
type MapVertex map[int64]*Vertex

var n int64
var currentVertex *Vertex

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var e error
	mapVertex := make(MapVertex)
	if scanner.Scan() {
		if n, e = strconv.ParseInt(scanner.Text(), 10, 64); e != nil {
			panic(e)
		}
	}
	for i := int64(0); scanner.Scan() && i < n; i++ {
		line := scanner.Text()
		linevals := strings.Split(line, " ")
		color := ColorRed
		if linevals[2] == "b" {
			color = ColorBlue
		}
		edge := &Edge{Color: color}

		var vertexNumber1, vertexNumber2 int64
		if vertexNumber1, e = strconv.ParseInt(linevals[0], 10, 64); e != nil {
			panic(e)
		}
		if vertexNumber2, e = strconv.ParseInt(linevals[1], 10, 64); e != nil {
			panic(e)
		}
		var (
			vertex1 *Vertex
			vertex2 *Vertex
			ok      bool
		)
		if vertex1, ok = mapVertex[vertexNumber1]; !ok {
			vertex1 = &Vertex{
				ID:    vertexNumber1,
				Edges: []*Edge{edge},
			}
			mapVertex[vertexNumber1] = vertex1
		} else {
			vertex1.Edges = append(vertex1.Edges, edge)
		}
		if vertex2, ok = mapVertex[vertexNumber2]; !ok {
			vertex2 = &Vertex{
				ID:    vertexNumber2,
				Edges: []*Edge{edge},
			}
			mapVertex[vertexNumber2] = vertex2
		} else {
			vertex2.Edges = append(vertex2.Edges, edge)
		}
		edge.Vertex[0] = vertex1
		edge.Vertex[1] = vertex2
	}

	for i, v := range mapVertex {
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Printf("map id: %d\nmem loc: %p\nval:\n%+v\n", i, v, v)
		for _, e := range v.Edges {
			fmt.Printf("-------------------------\n")
			fmt.Printf("edges vertex: \n%+v\n", e.Vertex)
		}
	}
}
