package ticket_seller

import (
	"ObjectOrientedProgramming/procedural_programming/ticket_office"
)

type TicketSeller struct {
	ticketOffice *ticket_office.TicketOffice
}

func NewTicketSeller(ticketOffice *ticket_office.TicketOffice) *TicketSeller {
	return &TicketSeller{
		ticketOffice: ticketOffice,
	}
}

func (t *TicketSeller) TicketOffice() *ticket_office.TicketOffice {
	return t.ticketOffice
}
