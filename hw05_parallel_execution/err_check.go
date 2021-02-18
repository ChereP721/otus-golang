package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type errCounter struct {
	max    int // max=0 - игнорируем ошибки
	cnt    int
	doneCh chan<- struct{}
	mu     sync.Mutex
	err    error
}

func newErrCounter(max int, doneCh chan<- struct{}) errCounter {
	return errCounter{
		max:    max,
		mu:     sync.Mutex{},
		doneCh: doneCh,
	}
}

func (ec *errCounter) inc() bool {
	ec.mu.Lock()
	defer ec.mu.Unlock()

	ec.cnt++

	if ec.cnt == ec.max {
		close(ec.doneCh)
		ec.err = ErrErrorsLimitExceeded
		return true
	}

	return false
}

func (ec *errCounter) check(err error) bool {
	if err == nil {
		return false
	}

	return ec.inc()
}
