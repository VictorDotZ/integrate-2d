package integrate

import (
	"integrate2d/pkg/points"
	"integrate2d/pkg/rules"
)

type TriangleVertexIds struct {
	Aidx uint64
	Bidx uint64
	Cidx uint64
}

func Integrate(vertices []points.Point2d, vertexIds []TriangleVertexIds, f func(points.Point2d) float64) float64 {
	result := float64(0)
	for _, vertex := range vertexIds {
		A := vertices[vertex.Aidx]
		B := vertices[vertex.Bidx]
		C := vertices[vertex.Cidx]
		result += rules.TriangleQuadrature(A, B, C, f)
	}
	return result
}
