package paxos

import "fmt"

type Proposer struct {
	id        int
	round     int
	number    int
	acceptors []int
}

func (p *Proposer) proposalNumber() int {
	return p.round<<16 | p.id
}

func (p *Proposer) majority() int {
	return len(p.acceptors)/2 + 1
}

func (p *Proposer) propose(v interface{}) interface{} {
	p.round++
	p.number = p.proposalNumber()

	prepareCount := 0
	maxNumber := 0
	for _, aid := range p.acceptors {
		args := MsgArgs{
			Number: p.number,
			From:   p.id,
			To:     aid,
		}
		reply := new(MsgReply)
		err := call(fmt.Sprintf("127.0.0.1:%d", aid), "Acceptor.Prepare", args, reply)
		if !err {
			continue
		}

		if reply.Ok {
			prepareCount++
			if reply.Number > maxNumber {
				maxNumber = reply.Number
				v = reply.Value
			}
		}
	}

	acceptCount := 0
	if prepareCount == p.majority() {
		for _, aid := range p.acceptors {
			args := MsgArgs{
				Number: p.number,
				Value:  v,
				From:   p.id,
				To:     aid,
			}
			reply := new(MsgReply)
			ok := call(fmt.Sprintf("127.0.0.1:%d", aid), "Acceptor.Accept", args, reply)
			if !ok {
				continue
			}

			if reply.Ok {
				acceptCount++
			}
		}
	}

	if acceptCount >= p.majority() {
		return v
	}

	return nil
}
