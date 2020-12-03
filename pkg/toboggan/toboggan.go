package toboggan

//TreeMap shows where the Toboggan can go
type TreeMap struct {
	depth int
	width int
	lines []string
}

//NewTreeMap creates a new map of the trees
func NewTreeMap(lines []string) *TreeMap {
	depth := len(lines)
	width := len(lines[0])
	return &TreeMap{depth: depth, width: width, lines: lines}
}

//Traverse with the toboggan
func (t *TreeMap) Traverse(x, y int) int {
	xpos := 0
	ypos := 0

	const tree byte = '#'

	treeCount := 0

	for {
		xpos = xpos + x
		ypos = ypos + y
		if xpos >= t.width {
			xpos = xpos - t.width
		}
		if ypos >= t.depth {
			break
		}
		if t.lines[ypos][xpos] == tree {
			treeCount++
		}
	}
	return treeCount
}
