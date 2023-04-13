package triangulation

import (
	"fmt"
	"integrate2d/pkg/integrate"
	"integrate2d/pkg/points"
	"io"
)

func GenerateTriangulation(Lx, Ly float64, Nx, Ny uint64, out chan string) {
	out <- fmt.Sprintf("%d\n", (Nx+1)*(Ny+1))
	out <- fmt.Sprintf("%d\n", 2*Nx*Ny)
	out <- fmt.Sprintf("%d\n", 3*Nx*Ny-Nx-Ny)
	out <- fmt.Sprintf("%d\n", 2*Nx*2*Ny)

	GenerateVetrex(Lx, Ly, Nx, Ny, out)
	GenerateTriangles(Nx, Ny, out)
	GenerateInnerEdges(Nx, Ny, out)
	GenerateOuterEdges(Nx, Ny, out)

	close(out)
}

func GenerateVetrex(Lx, Ly float64, Nx, Ny uint64, out chan string) {
	xLength := Lx / float64(Nx)
	yLength := Ly / float64(Ny)
	vertexIdx := 0
	for i := uint64(0); i <= Nx; i++ {
		for j := uint64(0); j <= Ny; j++ {
			out <- fmt.Sprintf("%d %e %e\n", vertexIdx, xLength*float64(i), yLength*float64(j))
			vertexIdx++
		}
	}
}

func GenerateTriangles(Nx, Ny uint64, out chan string) {
	triangleIdx := 0
	for i := uint64(1); i <= Nx; i++ {
		for j := uint64(1); j <= Ny; j++ {
			out <- fmt.Sprintf("%d %d %d %d\n",
				triangleIdx,
				(Ny+1)*(i-1)+(j-1),
				(Ny+1)*(i-1)+j,
				(Ny+1)*i+(j-1))
			triangleIdx++

			out <- fmt.Sprintf("%d %d %d %d\n", triangleIdx,
				(Ny+1)*(i-1)+j,
				(Ny+1)*i+j-1,
				(Ny+1)*i+j)
			triangleIdx++
		}
	}
}

func GenerateInnerEdges(Nx, Ny uint64, out chan string) {
	inner_edge_number := 0
	for i := uint64(1); i <= Nx; i++ {
		for j := uint64(1); j <= Ny; j++ {
			out <- fmt.Sprintf("%d %d %d\n", inner_edge_number,
				(Ny+1)*(i-1)+j,
				(Ny+1)*i+j-1)
			inner_edge_number++
			if i != Nx {
				out <- fmt.Sprintf("%d %d %d\n", inner_edge_number,
					(Ny+1)*i+j,
					(Ny+1)*i+j-1)
			}
			inner_edge_number++
			if j != Ny {
				out <- fmt.Sprintf("%d %d %d\n", inner_edge_number,
					(Ny+1)*(i-1)+j,
					(Ny+1)*i+j)
			}
			inner_edge_number++
		}
	}
}

func GenerateOuterEdges(Nx, Ny uint64, out chan string) {
	outer_edge_number := 0
	for i := uint64(0); i < Nx; i++ {
		out <- fmt.Sprintf("%d %d %d\n",
			outer_edge_number,
			(Ny+1)*i,
			(Ny+1)*(i+1))
		outer_edge_number++
		out <- fmt.Sprintf("%d %d %d\n",
			outer_edge_number,
			(Ny+1)*i+Ny,
			(Ny+1)*(i+1)+Ny)
		outer_edge_number++
	}

	for j := uint64(0); j < Ny; j++ {
		out <- fmt.Sprintf("%d %d %d\n",
			outer_edge_number,
			j+0,
			j+1)
		outer_edge_number++
		out <- fmt.Sprintf("%d %d %d\n",
			outer_edge_number,
			(Ny+1)*Nx+j+0,
			(Ny+1)*Nx+j+1)
		outer_edge_number++
	}
}

func ReadTriangleVertexIds(in io.Reader, numVertices, numTriangles uint64) (*[]integrate.TriangleVertexIds, error) {
	vertexIds := make([]integrate.TriangleVertexIds, 0)
	var current integrate.TriangleVertexIds
	var dummy uint64
	for i := uint64(0); i < numTriangles; i++ {
		current = integrate.TriangleVertexIds{}

		if n, err := fmt.Fscanf(in, "%d %d %d %d", &dummy, &current.Aidx, &current.Bidx, &current.Cidx); n != 4 || err != nil {
			return nil, fmt.Errorf("")
		}

		if current.Aidx >= numVertices || current.Bidx >= numVertices || current.Cidx >= numVertices {
			return nil, fmt.Errorf("")
		}

		vertexIds = append(vertexIds, current)
	}

	return &vertexIds, nil
}

func ReadVertices(in io.Reader, numVertices uint64) (*[]points.Point2d, error) {
	var current points.Point2d
	var dummy uint64
	vertices := make([]points.Point2d, 0)

	for i := uint64(0); i < numVertices; i++ {
		current = points.Point2d{}
		if n, err := fmt.Fscanf(in, "%d %e %e", &dummy, &current.X, &current.Y); n != 3 || err != nil {
			return nil, fmt.Errorf("")
		}

		vertices = append(vertices, current)
	}
	return &vertices, nil
}
