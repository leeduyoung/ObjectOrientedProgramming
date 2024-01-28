package bag

import (
	"ObjectOrientedProgramming/object_orientied_programming/invitation"
	"ObjectOrientedProgramming/object_orientied_programming/ticket"
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

func (b *Bag) Hold(ticket *ticket.Ticket) float64 {
	if b.HasInvitation() {
		log.Println("audience has invitation.")

		b.SetTicket(ticket)
		return 0
	} else {
		log.Println("audi no invitation.")

		b.MinusAmount(ticket.Fee())
		b.SetTicket(ticket)
		return ticket.Fee()
	}
}
