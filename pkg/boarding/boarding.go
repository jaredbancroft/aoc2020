package boarding

import (
	"strconv"
	"strings"
)

//Pass is a boarding pass
type Pass struct {
	Row, Column, SeatID int
}

//NewPass creates a new boarding pass
func NewPass(code string) Pass {
	bin := func(r rune) rune {
		if r == 'B' || r == 'R' {
			return '1'
		}
		return '0'

	}

	row, _ := strconv.ParseInt(strings.Map(bin, code[:7]), 2, 64)
	column, _ := strconv.ParseInt(strings.Map(bin, code[7:]), 2, 64)
	seatID := int(row*8 + column)

	return Pass{Row: int(row), Column: int(column), SeatID: seatID}
}

//BySeatID custom type to implement the sort.Interface
type BySeatID []Pass

//Len is needed to implement sort.Interface
func (s BySeatID) Len() int {
	return len(s)
}

//Swap is needed to implement sort.Interface
func (s BySeatID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//Less is needed to implement sort.Interface
func (s BySeatID) Less(i, j int) bool {
	return s[i].SeatID < s[j].SeatID
}
