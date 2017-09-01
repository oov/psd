package composite

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var ErrAborted = errors.New("composite: aborted")

type parallelContext struct {
	M     sync.RWMutex
	Wg    sync.WaitGroup
	abort int32
}

func (pc *parallelContext) Wait(ctx context.Context) error {
	complete := make(chan struct{})
	go func() {
		pc.Wg.Wait()
		complete <- struct{}{}
	}()
	select {
	case <-complete:
		return nil
	case <-ctx.Done():
		pc.SetAbort()
		<-complete
		return ErrAborted
	}
}

func (pc *parallelContext) SetAbort() {
	if pc == nil {
		return
	}
	atomic.StoreInt32(&pc.abort, 1)
}

func (pc *parallelContext) Aborted() bool {
	if pc == nil {
		return false
	}
	return atomic.LoadInt32(&pc.abort) != 0
}

func (pc *parallelContext) Done() {
	if pc == nil {
		return
	}
	pc.Wg.Done()
}
