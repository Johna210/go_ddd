package chapter3

// two value objects with the same value are the same
// the fields should not be exported and mutated
// value objects should be replaceable

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

const (
	directionUnknown = iota
	directionNorth
	directionSouth
	directionEast
	directionWest
)

func TrackPlayer() {
	currLocation := NewPoint(3, 4)
	currLocation = move(currLocation, directionNorth)
}

func move(currLocation Point, direction int) Point {
	switch direction {
	case directionNorth:
		return NewPoint(currLocation.x, currLocation.y+1)
	case directionSouth:
		return NewPoint(currLocation.x, currLocation.y-1)
	case directionEast:
		return NewPoint(currLocation.x+1, currLocation.y)
	case directionWest:
		return NewPoint(currLocation.x-1, currLocation.y)
	default:
		// do something
	}
	return currLocation
}
