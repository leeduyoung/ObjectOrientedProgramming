package main

import (
	"ObjectOrientedProgramming/procedural_programming/audience"
	"ObjectOrientedProgramming/procedural_programming/bag"
	"ObjectOrientedProgramming/procedural_programming/invitation"
	"ObjectOrientedProgramming/procedural_programming/theater"
	"ObjectOrientedProgramming/procedural_programming/ticket"
	"ObjectOrientedProgramming/procedural_programming/ticket_office"
	"ObjectOrientedProgramming/procedural_programming/ticket_seller"
	"time"
)

func main() {

	var tickets []*ticket.Ticket
	for i := 0; i < 10; i++ {
		tickets = append(tickets, ticket.NewTicket(10000))
	}

	var audiences = []*audience.Audience{
		audience.NewAudience(
			bag.NewBag(50000, &invitation.Invitation{When: time.Now().Add(time.Hour * 3)}),
		),
		audience.NewAudience(
			bag.NewBag(50000, nil),
		),
	}

	for _, v := range audiences {
		newTheater := theater.NewTheater(
			ticket_seller.NewTicketSeller(
				ticket_office.NewTicketOffice(0, tickets),
			),
		)

		newTheater.Enter(v)
	}
}
