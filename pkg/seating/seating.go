package seating

import (
	"errors"
	"fmt"
)

//SeatState is a custom type to map the different states
type SeatState int

const (
	empty SeatState = iota
	occupied
	floor
)

//Seat represents a seat which has 1 of 3 states
type Seat struct {
	State SeatState
}

func (s *Seat) diffSeat(seat Seat) bool {
	if s.State == seat.State {
		return true
	}
	return false
}

//Toggle toggles a seat from empty to occupied, etc
func (s *Seat) toggle() {
	if s.State == empty {
		s.State = occupied
	} else if s.State == occupied {
		s.State = empty
	}
}

//NewSeat creates a new seat
func NewSeat(state SeatState) Seat {
	return Seat{State: state}
}

//SeatLayout is a floor mapping of the seats
type SeatLayout struct {
	Seats [][]Seat
}

//NewSeatLayout initializes a seat layout
func NewSeatLayout(layout []string) (*SeatLayout, error) {
	parsedSeats := [][]Seat{}
	for _, seats := range layout {
		parsedRow := []Seat{}
		for _, seat := range seats {
			switch seat {
			case 'L':
				parsedRow = append(parsedRow, NewSeat(empty))
			case '.':
				parsedRow = append(parsedRow, NewSeat(floor))
			case '#':
				return nil, errors.New("Invalid input")
			default:
				return nil, errors.New("Invalid input")
			}
		}
		parsedSeats = append(parsedSeats, parsedRow)
	}
	return &SeatLayout{Seats: parsedSeats}, nil
}

func (l *SeatLayout) same(previous [][]Seat) bool {
	for i := 0; i < len(l.Seats[0]); i++ {
		for j := 0; j < len(l.Seats); j++ {
			if l.Seats[j][i] != previous[j][i] {
				return false
			}
		}
	}
	return true
}

//Count returns the number of occupied seats
func (l *SeatLayout) Count() int {
	count := 0
	for _, ypos := range l.Seats {
		for _, xpos := range ypos {
			if xpos.State == occupied {
				count++
			}
		}
	}
	return count
}

func (l *SeatLayout) deepCopy() [][]Seat {
	previous := [][]Seat{}
	for _, ypos := range l.Seats {
		row := []Seat{}
		for _, xpos := range ypos {
			row = append(row, NewSeat(xpos.State))
		}
		previous = append(previous, row)
	}
	return previous
}

//Model runs the model rules against the current seat layout
func (l *SeatLayout) Model() {
	for {
		previous := l.deepCopy()

		toChange := [][]int{}

		for j, ypos := range l.Seats {
			for i, xpos := range ypos {
				if l.checkAdjacent(i, j, xpos.State) {
					toChange = append(toChange, []int{i, j})
				}
			}
		}

		for _, seat := range toChange {
			l.Seats[seat[1]][seat[0]].toggle()
		}

		if l.same(previous) {
			break
		}
	}
}

//Model2 models part2
func (l *SeatLayout) Model2() {
	for {
		previous := l.deepCopy()
		toChange := [][]int{}
		for j, ypos := range l.Seats {
			for i, xpos := range ypos {
				if l.checkAdjacentLOS(i, j, xpos.State) {
					toChange = append(toChange, []int{i, j})
				}
			}
		}

		for _, seat := range toChange {
			l.Seats[seat[1]][seat[0]].toggle()
		}
		if l.same(previous) {
			break
		}
	}
}

func (l *SeatLayout) checkAdjacentLOS(x, y int, state SeatState) bool {
	occupiedCount := 0

	//Down Right
	i := x
	j := y
	for {
		i++
		j++
		if i < len(l.Seats[0]) && j < len(l.Seats) {
			if l.Seats[j][i].State != floor {
				if l.Seats[j][i].State == occupied {
					occupiedCount++
					break
				}
				break
			}
		} else {
			break
		}
	}
	//Down Mid
	i = x
	j = y
	for {
		j++
		if j < len(l.Seats) {
			if l.Seats[j][i].State != floor {
				if l.Seats[j][i].State == occupied {
					occupiedCount++
					break
				}
				break
			}
		} else {
			break
		}
	}
	//Down Left
	i = x
	j = y
	for {
		i--
		j++
		if i >= 0 && j < len(l.Seats) {
			if l.Seats[j][i].State != floor {
				if l.Seats[j][i].State == occupied {
					occupiedCount++
					break
				}
				break
			}
		} else {
			break
		}
	}
	//Left
	i = x
	j = y
	for {
		i--
		if i >= 0 {
			if l.Seats[j][i].State != floor {
				if l.Seats[j][i].State == occupied {
					occupiedCount++
					break
				}
				break
			}
		} else {
			break
		}
	}
	//Right
	i = x
	j = y
	for {
		i++
		if i < len(l.Seats[0]) {
			if l.Seats[j][i].State != floor {
				if l.Seats[j][i].State == occupied {
					occupiedCount++
					break
				}
				break
			}
		} else {
			break
		}
	}
	//Up Right
	i = x
	j = y
	for {
		i++
		j--
		if i < len(l.Seats[0]) && j >= 0 {
			if l.Seats[j][i].State != floor {
				if l.Seats[j][i].State == occupied {
					occupiedCount++
					break
				}
				break
			}
		} else {
			break
		}
	}
	//Up Mid
	i = x
	j = y
	for {
		j--
		if j >= 0 {
			if l.Seats[j][i].State != floor {
				if l.Seats[j][i].State == occupied {
					occupiedCount++
					break
				}
				break
			}
		} else {
			break
		}
	}
	//Up Left
	i = x
	j = y
	for {
		i--
		j--
		if i >= 0 && j >= 0 {
			if l.Seats[j][i].State != floor {
				if l.Seats[j][i].State == occupied {
					occupiedCount++
					break
				}
				break
			}
		} else {
			break
		}
	}

	if (occupiedCount >= 5 && state == occupied) || (occupiedCount == 0 && state == empty) {
		return true
	}
	return false
}

func (l *SeatLayout) checkAdjacent(i, j int, state SeatState) bool {
	occupiedCount := 0
	UL := []int{i - 1, j - 1}
	UM := []int{i, j - 1}
	UR := []int{i + 1, j - 1}
	ML := []int{i - 1, j}
	MR := []int{i + 1, j}
	LL := []int{i - 1, j + 1}
	LM := []int{i, j + 1}
	LR := []int{i + 1, j + 1}

	adjacencies := [][]int{}
	adjacencies = append(adjacencies, UL, UM, UR, ML, MR, LL, LM, LR)

	for _, adj := range adjacencies {
		if adj[0] >= 0 && adj[1] >= 0 && adj[0] < len(l.Seats[0]) && adj[1] < len(l.Seats) {
			if l.Seats[adj[1]][adj[0]].State == occupied {
				occupiedCount++
			}
		}
	}

	if (occupiedCount >= 4 && state == occupied) || (occupiedCount == 0 && state == empty) {
		return true
	}

	return false
}

//PrettyPrint the floor layout
func (l *SeatLayout) PrettyPrint() {
	for _, ypos := range l.Seats {
		for _, xpos := range ypos {
			if xpos.State == occupied {
				fmt.Print("#")
			}
			if xpos.State == empty {
				fmt.Print("L")
			}
			if xpos.State == floor {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
