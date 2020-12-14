package docking

import (
	"strconv"
)

//Mask holds the bits that need to be changed in a value
type Mask struct {
	Bits map[int]int
	Xs   []int
}

//NewMask creates a new mask
func NewMask(mask string) Mask {
	bitmap := make(map[int]int)
	xmap := make([]int, 0)
	for i := 0; i < len(mask); i++ {
		if mask[i] != 'X' {
			bitvalue, _ := strconv.Atoi(string(mask[i]))
			bitmap[len(mask)-i-1] = bitvalue
		}
		if mask[i] == 'X' {
			xmap = append(xmap, len(mask)-i-1)
		}
	}
	return Mask{Bits: bitmap, Xs: xmap}
}

//SetBit will set the bit at a position to 1
func (m Mask) SetBit(n int, pos int) int {
	n |= (1 << pos)
	return n
}

//ClearBit will set the bit at a position to 0
func (m Mask) ClearBit(n int, pos int) int {
	mask := ^(1 << pos)
	n &= mask
	return n
}

//MemMask is a mask for memory addresses
type MemMask struct {
	MemAddrs map[int]int
}

//NewMemMask creates a new memory mask for a given address
func NewMemMask(memAddr int, mask Mask, value int) MemMask {
	memAddrs := make(map[int]int)
	for pos, bit := range mask.Bits {
		if bit == 1 {
			memAddr = mask.SetBit(memAddr, pos)
		}
	}

	mask.recurse(mask.Xs, memAddr, memAddrs, value)

	return MemMask{MemAddrs: memAddrs}
}

func (m Mask) recurse(xs []int, memAddr int, memAddrs map[int]int, value int) {

	if len(xs) > 0 {
		low := m.ClearBit(memAddr, xs[0])
		if _, ok := memAddrs[low]; !ok {
			memAddrs[low] = value
		}
		high := m.SetBit(memAddr, xs[0])
		if _, ok := memAddrs[high]; !ok {
			memAddrs[high] = value
		}
		m.recurse(xs[1:], low, memAddrs, value)
		m.recurse(xs[1:], high, memAddrs, value)
	}
}
