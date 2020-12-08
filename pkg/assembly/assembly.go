package assembly

//Instruction is a single instruction comprised of an operation and an argument
type Instruction struct {
	Op      string
	Arg     int
	visited bool
}

//NewInstruction creates a new instruction instance
func NewInstruction(operation string, argument int) Instruction {
	return Instruction{Op: operation, Arg: argument, visited: false}
}

func (i *Instruction) visit() {
	i.visited = true
}

func (i *Instruction) unvisit() {
	i.visited = false
}

func (i *Instruction) swap() {
	switch i.Op {
	case "jmp":
		i.Op = "nop"
	case "nop":
		i.Op = "jmp"
	}
}

//InstructionList is a slice of instructions
type InstructionList struct {
	Instructions []Instruction
}

//NewInstructionList initialiazes a new empty list of instructions
func NewInstructionList() *InstructionList {
	return &InstructionList{}
}

//Append appends a new instruction to the list
func (il *InstructionList) Append(instruction Instruction) {
	il.Instructions = append(il.Instructions, instruction)
}

func (il *InstructionList) resetAllVisited() {
	for i := 0; i < len(il.Instructions); i++ {
		il.Instructions[i].unvisit()
	}
}

//Process the list of instructions
func (il *InstructionList) Process() int {
	acc := 0
	i := 0

	for {
		if il.Instructions[i].visited == false {
			il.Instructions[i].visit()
			switch il.Instructions[i].Op {
			case "jmp":
				i = i + il.Instructions[i].Arg
			case "acc":
				acc = acc + il.Instructions[i].Arg
				i++
			case "nop":
				i++
			}
		} else {
			break
		}
	}

	return acc
}

//Fix fixes the program
func (il *InstructionList) Fix() int {
	acc := 0
	i := 0

	il.resetAllVisited()

	for j := 0; j < len(il.Instructions); j++ {

		if il.Instructions[j].Op == "nop" || il.Instructions[j].Op == "jmp" {
			il.Instructions[j].swap()
		}

		for {
			if i == len(il.Instructions) {
				return acc
			}

			if il.Instructions[i].visited == false {
				il.Instructions[i].visit()
				switch il.Instructions[i].Op {
				case "jmp":
					i = i + il.Instructions[i].Arg
				case "acc":
					acc = acc + il.Instructions[i].Arg
					i++
				case "nop":
					i++
				}
			} else {
				break
			}
		}

		if il.Instructions[j].Op == "nop" || il.Instructions[j].Op == "jmp" {
			il.Instructions[j].swap()
		}
		il.resetAllVisited()
		i = 0
		acc = 0
	}

	return acc
}
