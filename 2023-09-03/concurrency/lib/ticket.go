package lib

type GoTickets interface {
	Take()
	Return()
	Active() bool
	Total() uint32
	Remainder() uint32
}

type myTicket struct {
	total    uint32
	ticketCh chan struct{}
	active   bool
}

func (m *myTicket) Take() {
	//TODO implement me
	panic("implement me")
}

func (m *myTicket) Return() {
	//TODO implement me
	panic("implement me")
}

func (m *myTicket) Active() bool {
	return m.active
}

func (m *myTicket) Total() uint32 {
	return m.total
}

func (m *myTicket) Remainder() uint32 {
	//TODO implement me
	panic("implement me")
}
