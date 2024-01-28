package bag

import (
	"ObjectOrientedProgramming/procedural_programming/invitation"
	"ObjectOrientedProgramming/procedural_programming/ticket"
	"log"
)

type Bag struct {
	amount     float64
	invitation *invitation.Invitation
	ticket     *ticket.Ticket
}

func NewBag(amount float64, invitation *invitation.Invitation) *Bag {
	return &Bag{
		amount:     amount,
		invitation: invitation,
		ticket:     nil,
	}
}

func (b *Bag) HasInvitation() bool {
	return b.invitation != nil
}

func (b *Bag) HasTicket() bool {
	return b.ticket != nil
}

func (b *Bag) SetTicket(ticket *ticket.Ticket) {
	log.Println("audience set ticket")

	b.ticket = ticket
}

func (b *Bag) MinusAmount(amount float64) {
	log.Println("audience minus amount.", amount)

	b.amount -= amount
}

func (b *Bag) PlusAmount(amount float64) {
	log.Println("audience plus amount.", amount)

	b.amount += amount
}
