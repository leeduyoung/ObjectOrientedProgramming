package ticket

type Ticket struct {
	fee float64
}

func (t *Ticket) Fee() float64 {
	return t.fee
}
