package ticket_office

import (
	"ObjectOrientedProgramming/object_orientied_programming/audience"
	"ObjectOrientedProgramming/object_orientied_programming/ticket"
	"log"
)

type TicketOffice struct {
	amount  float64
	tickets []*ticket.Ticket
}

func NewTicketOffice(amount float64, tickets []*ticket.Ticket) *TicketOffice {
	return &TicketOffice{
		amount:  amount,
		tickets: tickets,
	}
}

func (t *TicketOffice) Ticket() *ticket.Ticket {
	log.Println("ticket office get ticket")

	firstTicket := t.tickets[0]
	t.tickets = t.tickets[1:]
	return firstTicket
}

func (t *TicketOffice) PlusAmount(amount float64) {
	log.Println("ticket office plus amount.", amount)

	t.amount += amount
}

func (t *TicketOffice) MinusAmount(amount float64) {
	log.Println("ticket office minus amount.", amount)

	t.amount -= amount
}

func (t *TicketOffice) SellTicketTo(audi *audience.Audience) {
	t.PlusAmount(audi.Buy(t.Ticket()))
}
