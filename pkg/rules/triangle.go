package rules

import (
	"integrate2d/pkg/areas"
	"integrate2d/pkg/points"
)

func TriangleQuadrature(A, B, C points.Point2d, f func(points.Point2d) float64) float64 {
	return (areas.GetTriangleArea(A, B, C) / 3.0) * (f(points.GetMiddlePoint(A, B)) + f(points.GetMiddlePoint(B, C)) + f(points.GetMiddlePoint(C, A)))
}
