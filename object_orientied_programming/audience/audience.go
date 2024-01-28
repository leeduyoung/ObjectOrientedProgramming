package audience

import (
	"ObjectOrientedProgramming/object_orientied_programming/bag"
	"ObjectOrientedProgramming/object_orientied_programming/ticket"
)

type Audience struct {
	bag *bag.Bag
}

func NewAudience(bag *bag.Bag) *Audience {
	return &Audience{
		bag: bag,
	}
}

func (a *Audience) Bag() *bag.Bag {
	return a.bag
}

func (a *Audience) Buy(ticket *ticket.Ticket) float64 {
	return a.bag.Hold(ticket)
}
