package concurrency

import (
	"context"
	"learn365/2023-09-03/concurrency/lib"
	"sync/atomic"
)

type Generator interface {
	Start()
	Stop() bool
	Status() uint32
	CallCount() int64
}

type generator struct {
	isStart       atomic.Bool
	caller        lib.Caller
	concurrency   uint32
	runningStatus uint32
	ctx           context.Context
	cancel        context.CancelFunc
	callCount     uint32
	resultCh      chan lib.CallResult
}

func (g *generator) Start() {
	//TODO implement me
	panic("implement me")
}

func (g *generator) Stop() bool {
	//TODO implement me
	panic("implement me")
}

func (g *generator) Status() uint32 {
	return atomic.LoadUint32(&g.runningStatus)
}

func (g *generator) CallCount() uint32 {
	return atomic.LoadUint32(&g.callCount)
}
