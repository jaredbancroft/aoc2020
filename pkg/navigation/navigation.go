package navigation

import (
	"errors"
	"math"
)

//Position holds the current location
type Position struct {
	X, Y int
}

//ManhattanDistance returns the manhattan distance of the location
func (p *Position) ManhattanDistance() int {
	return int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
}

//NewPosition creates and initializes a new location with x and y coordinates
func NewPosition(x, y int) *Position {
	return &Position{X: x, Y: y}
}

//Direction is a cardinal direction interface
type Direction interface {
	GetLetter() string
	GetValue() int
}

//North implements Direction interface
type North struct{}

//GetLetter gets the short code for the direction
func (n North) GetLetter() string {
	return "N"
}

//GetValue gets the int code for the direction
func (n North) GetValue() int {
	return 0
}

//East implements Direction interface
type East struct{}

//GetLetter gets the short code for the direction
func (e East) GetLetter() string {
	return "E"
}

//GetValue gets the int code for the direction
func (e East) GetValue() int {
	return 1
}

//South implements Direction interface
type South struct{}

//GetLetter gets the short code for the direction
func (s South) GetLetter() string {
	return "S"
}

//GetValue gets the int code for the direction
func (s South) GetValue() int {
	return 2
}

//West implements Direction interface
type West struct{}

//GetLetter gets the short code for the direction
func (w West) GetLetter() string {
	return "W"
}

//GetValue gets the int code for the direction
func (w West) GetValue() int {
	return 3
}

//NewDirection returns a new cardinal direction
func NewDirection(value int) (Direction, error) {
	switch value {
	case 0:
		return North{}, nil
	case 1:
		return East{}, nil
	case 2:
		return South{}, nil
	case 3:
		return West{}, nil
	default:
		return nil, errors.New("No such direction")
	}
}

//Boat has a Heading and a Posi
type Boat struct {
	Heading Direction
	Position
}

//Boat2 is a boat with a waypoint
type Boat2 struct {
	Position
	Waypoint
}

//NewBoat gets you a new boat
func NewBoat(direction Direction, position Position) *Boat {
	return &Boat{Heading: direction, Position: position}
}

//NewBoat2 gets you a new boat
func NewBoat2(position Position, waypoint Waypoint) *Boat2 {
	return &Boat2{Position: position, Waypoint: waypoint}
}

//Move will move the boat according to the action and the value
func (b *Boat) Move(action string, value int) {
	switch action {
	case "N":
		b.Position.Y = b.Position.Y + value
	case "S":
		b.Position.Y = b.Position.Y - value
	case "E":
		b.Position.X = b.Position.X + value
	case "W":
		b.Position.X = b.Position.X - value
	case "R":
		b.turn(action, value)
	case "L":
		b.turn(action, value)
	case "F":
		b.Move(b.Heading.GetLetter(), value)
	}
}

func (b *Boat) turn(action string, value int) {
	numTurns := value / 90
	switch action {
	case "R":
		newHeading := b.Heading.GetValue() + numTurns
		if newHeading >= 4 {
			newHeading = newHeading - 4
		}
		b.Heading, _ = NewDirection(newHeading)
	case "L":
		newHeading := b.Heading.GetValue() - numTurns
		if newHeading < 0 {
			newHeading = newHeading + 4
		}
		b.Heading, _ = NewDirection(newHeading)
	}

}

//Move the boat2
func (b *Boat2) Move(action string, value int) {

	switch action {
	case "N":
		b.Waypoint.Position.Y = b.Waypoint.Position.Y + value

	case "S":
		b.Waypoint.Position.Y = b.Waypoint.Position.Y - value

	case "E":
		b.Waypoint.Position.X = b.Waypoint.Position.X + value

	case "W":
		b.Waypoint.Position.X = b.Waypoint.Position.X - value

	case "R":
		b.turn(action, value)

	case "L":
		b.turn(action, value)

	case "F":
		shipX := value * b.Waypoint.Position.X
		shipY := value * b.Waypoint.Position.Y
		b.Position.X = b.Position.X + shipX
		b.Position.Y = b.Position.Y + shipY
	}
}

func (b *Boat2) turn(action string, value int) {
	numTurns := value / 90
	switch action {
	case "R":
		newX := 0
		newY := 0
		for i := 0; i < numTurns; i++ {
			newY = -b.Waypoint.Position.X
			newX = b.Waypoint.Position.Y

			b.Waypoint.Position.X = newX
			b.Waypoint.Position.Y = newY
		}

	case "L":
		newX := 0
		newY := 0
		for i := 0; i < numTurns; i++ {
			newY = b.Waypoint.Position.X
			newX = -b.Waypoint.Position.Y

			b.Waypoint.Position.X = newX
			b.Waypoint.Position.Y = newY
		}

	}
}

//Waypoint the boat is trying to reach
type Waypoint struct {
	Position
}

//NewWaypoint creates a new waypoint
func NewWaypoint(position Position) *Waypoint {
	return &Waypoint{Position: position}
}
