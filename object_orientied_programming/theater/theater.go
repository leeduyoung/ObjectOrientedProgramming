package theater

import (
	"ObjectOrientedProgramming/object_orientied_programming/audience"
	"ObjectOrientedProgramming/object_orientied_programming/ticket_seller"
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

func (t *Theater) Enter(audi *audience.Audience) {
	log.Println("audience is enter the theater.")

	t.ticketSeller.SellTo(audi)
}
