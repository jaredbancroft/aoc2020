package shuttle

import "fmt"

//Subject interface is implemented by concrete subjects
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObserver()
}

//Timer is a concrete subject
type Timer struct {
	time         int
	observerList []Observer
}

//Observer interface is implemented by concrete observers
type Observer interface {
	Update(value int)
}

//Bus is a concrete observer
type Bus struct {
	id            int
	departing     bool
	time          int
	simpleSubject Subject
}

//NewBus creates a new Bus
func NewBus(id int, ss Subject) *Bus {
	newBus := &Bus{
		id:            id,
		simpleSubject: ss,
	}
	newBus.simpleSubject.RegisterObserver(newBus)
	return newBus
}

//Display will display bus info
func (b *Bus) Display() {
	fmt.Printf("Bus %v is departing: %v at %v\n", b.id, b.departing, b.time)
}

//Update updates the buses time
func (b *Bus) Update(time int) {
	b.time = time
	if b.time%b.id == 0 {
		b.departing = true
	} else {
		b.departing = false
	}
	//b.Display()
}

//IsDeparting checks if a bus is departing
func (b *Bus) IsDeparting() bool {
	return b.departing
}

//GetID gets the buses ID
func (b *Bus) GetID() int {
	return b.id
}

//GetTime returns the current time observed
func (b *Bus) GetTime() int {
	return b.time
}

//NewTimer gives us an instance of a timer
func NewTimer() *Timer {
	return &Timer{
		time:         0,
		observerList: make([]Observer, 0),
	}
}

//RegisterObserver adds an observer to the subjects observer list
func (t *Timer) RegisterObserver(o Observer) {
	t.observerList = append(t.observerList, o)
}

//RemoveObserver removes an observer from the subjects observer list
func (t *Timer) RemoveObserver(o Observer) {
	found := false
	i := 0
	for ; i < len(t.observerList); i++ {
		if t.observerList[i] == o {
			found = true
			break
		}
	}
	if found {
		t.observerList = append(t.observerList[:i], t.observerList[i+1:]...)
	}
}

//NotifyObserver will notify of changes to the time
func (t *Timer) NotifyObserver() {
	for _, observer := range t.observerList {
		observer.Update(t.time)
	}
}

//SetTime updates the time on the timer
func (t *Timer) SetTime(time int) {
	t.time = time
	t.NotifyObserver()
}
