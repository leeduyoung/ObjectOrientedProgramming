package bag

import (
	"ObjectOrientedProgramming/invitation"
	"ObjectOrientedProgramming/ticket"
)

type Bag struct {
	amount     float64
	invitation *invitation.Invitation
	ticket     *ticket.Ticket
}

func (b *Bag) HasInvitation() bool {
	return b.invitation != nil
}

func (b *Bag) HasTicket() bool {
	return b.ticket != nil
}
