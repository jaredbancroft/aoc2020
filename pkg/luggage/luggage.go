package luggage

//Bag is a type of bag
type Bag struct {
	Name     string
	Contains map[string]int
}

//NewBag creates a new bag
func NewBag(name string) *Bag {
	return &Bag{Name: name}
}

//Walk walks the bags all the way down and gives you a total count
func (b *Bag) Walk(luggageList []Bag) int {
	count := 0
	for k, v := range b.Contains {
		kcount := 0
		for _, ll := range luggageList {
			if ll.Name == k {
				kcount = ll.Walk(luggageList)
			}
		}
		count = count + (v + v*kcount)
	}
	return count
}
