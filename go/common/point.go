package common

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p *Point) Add(t *Point) *Point {
	return &Point{X: p.X + t.X, Y: p.Y + t.Y}
}

// ManhattanDistance returns the distance between two points, defined as |x1-x2| + |y1-y2|
func (p *Point) ManhattanDistance(t *Point) int {
	return int(math.Abs(float64(p.X-t.X))) + int(math.Abs(float64(p.Y-t.Y)))
}

func (p *Point) Compare(t *Point) int {
	if p.X < t.X || p.Y < t.Y {
		return -1
	} else if p.X == t.X || p.Y == t.Y {
		return 0
	}

	return 1
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}
