package functions

import (
	"integrate2d/pkg/points"
	"math"
)

func Polynom(p points.Point2d) float64 {
	return math.Pow(p.X, 4) + math.Pow(p.X*p.Y, 2) + math.Pow(p.Y, 4)
}
