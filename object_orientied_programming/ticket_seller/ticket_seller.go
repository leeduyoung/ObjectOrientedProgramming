package ticket_seller

import (
	"ObjectOrientedProgramming/object_orientied_programming/audience"
	"ObjectOrientedProgramming/object_orientied_programming/ticket_office"
)

type TicketSeller struct {
	ticketOffice *ticket_office.TicketOffice
}

func NewTicketSeller(ticketOffice *ticket_office.TicketOffice) *TicketSeller {
	return &TicketSeller{
		ticketOffice: ticketOffice,
	}
}

func (t *TicketSeller) SellTo(audi *audience.Audience) {
	t.ticketOffice.SellTicketTo(audi)
}
