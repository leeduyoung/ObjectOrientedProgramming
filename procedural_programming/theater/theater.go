package theater

import (
	"ObjectOrientedProgramming/procedural_programming/audience"
	"ObjectOrientedProgramming/procedural_programming/ticket_seller"
	"log"
)

type Theater struct {
	ticketSeller *ticket_seller.TicketSeller
}

func NewTheater(seller *ticket_seller.TicketSeller) *Theater {
	return &Theater{
		ticketSeller: seller,
	}
}

func (t *Theater) Enter(audience *audience.Audience) {
	log.Println("audience is enter the theater.")

	if audience.Bag().HasInvitation() {
		log.Println("audience has invitation.")

		ticket := t.ticketSeller.TicketOffice().Ticket()
		audience.Bag().SetTicket(ticket)
	} else {
		log.Println("audience no invitation.")

		ticket := t.ticketSeller.TicketOffice().Ticket()
		audience.Bag().MinusAmount(ticket.Fee())
		t.ticketSeller.TicketOffice().PlusAmount(ticket.Fee())
		audience.Bag().SetTicket(ticket)
	}
}
