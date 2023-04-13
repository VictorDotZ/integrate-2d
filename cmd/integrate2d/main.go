package main

import (
	"fmt"
	"integrate2d/pkg/functions"
	"integrate2d/pkg/integrate"
	"integrate2d/pkg/triangulation"
	"os"
)

func main() {
	var dummy uint64
	var n_vertices uint64
	var n_triangles uint64

	if n, err := fmt.Fscanf(os.Stdin, "%d", &n_vertices); n != 1 || err != nil {
		os.Exit(1)
	}

	if n, err := fmt.Fscanf(os.Stdin, "%d", &n_triangles); n != 1 || err != nil {
		os.Exit(2)
	}

	if n, err := fmt.Fscanf(os.Stdin, "%d\n%d", &dummy, &dummy); n != 2 || err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	vertices, err := triangulation.ReadVertices(os.Stdin, n_vertices)

	if err != nil {
		os.Exit(4)
	}

	vertexIds, err := triangulation.ReadTriangleVertexIds(os.Stdin, n_vertices, n_triangles)

	if err != nil {

		fmt.Println(err)
		os.Exit(5)
	}

	result := integrate.Integrate(*vertices, *vertexIds, functions.Polynom)

	fmt.Printf("%e\n", result)
}
