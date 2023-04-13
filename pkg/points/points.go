package points

type Point2d struct {
	X float64
	Y float64
}

func GetMiddlePoint(l, r Point2d) Point2d {
	return Point2d{
		(l.X + r.X) / 2.0,
		(l.Y + r.Y) / 2.0,
	}
}
