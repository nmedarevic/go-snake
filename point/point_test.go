package point

import (
	"testing"
)

func TestDifferentPointsShouldNotBeSame(t *testing.T) {
	point1 := Point{X: 0, Y: 0}
	point2 := Point{X: 1, Y: 0}

	if point1.isEqual(&point2) {
		t.Errorf("Different Points should not be the same")
	}
}
