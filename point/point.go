package point

type Point struct {
	X uint8
	Y uint8
}

func (p *Point) isEqual(p2 *Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}
