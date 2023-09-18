package paxos

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Acceptor struct {
	lis            net.Listener
	id             int
	minProposal    int
	acceptedNumber int
	acceptedValue  interface{}
	learners       []int
}

func newAcceptor(id int, learners []int) *Acceptor {
	acceptor := &Acceptor{
		id:       id,
		learners: learners,
	}
	acceptor.server()
	return acceptor
}

func (a *Acceptor) server() {
	rpcServer := rpc.NewServer()
	_ = rpcServer.Register(a)
	addr := fmt.Sprintf(":%d", a.id)
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	a.lis = l

	go func() {
		for {
			conn, err := a.lis.Accept()
			if err != nil {
				continue
			}
			go rpcServer.ServeConn(conn)
		}
	}()
}

func (a *Acceptor) close() {
	_ = a.lis.Close()
}

func (a *Acceptor) Prepare(args *MsgArgs, reply *MsgReply) error {
	if args.Number > a.minProposal {
		a.minProposal = args.Number
		reply.Number = a.acceptedNumber
		reply.Value = a.acceptedValue
		reply.Ok = true
	} else {
		reply.Ok = false
	}

	return nil
}

func (a *Acceptor) Accept(args *MsgArgs, reply *MsgReply) error {
	if args.Number >= a.minProposal {
		a.minProposal = args.Number
		a.acceptedNumber = args.Number
		a.acceptedValue = args.Value
		reply.Ok = true
		for _, lid := range a.learners {
			go func(learner int) {
				addr := fmt.Sprintf("127.0.0.1:%d", learner)
				args.From = a.id
				args.To = learner
				resp := new(MsgReply)
				ok := call(addr, "Learner.Learn", args, resp)
				if !ok {
					return

				}
			}(lid)
		}
	} else {
		reply.Ok = false
	}
	return nil
}
