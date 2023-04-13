package areas

import (
	"integrate2d/pkg/points"
	"math"
)

func GetTriangleArea(A, B, C points.Point2d) float64 {
	return 0.5 * math.Abs((A.X*(B.Y-C.Y) + B.X*(C.Y-A.Y) + C.X*(A.Y-B.Y)))
}
