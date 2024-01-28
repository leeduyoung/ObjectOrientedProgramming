package ticket

type Ticket struct {
	fee float64
}

func NewTicket(fee float64) *Ticket {
	return &Ticket{
		fee: fee,
	}
}

func (t *Ticket) Fee() float64 {
	return t.fee
}
