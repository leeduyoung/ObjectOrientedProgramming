package audience

import "ObjectOrientedProgramming/procedural_programming/bag"

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
