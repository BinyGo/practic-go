// Synchronization Patterns
// semaphore
// Allows controlling access to a common resource
package semaphore

import (
	"errors"
	"os"
	"time"
)

var (
	ErrNoTickets      = errors.New("semaphore: could not acquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: cant't release the semaphore without acquiring if first")
)

type Interface interface {
	Acquire() error
	Release() error
}

type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *implementation) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return ErrNoTickets
	}
}

func (s *implementation) Release() error {
	select {
	case <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func New(tickets int, timeout time.Duration) Interface {
	return &implementation{
		sem:     make(chan struct{}, tickets),
		timeout: timeout,
	}
}

func Demo1() {
	tickets, timeout := 1, 30*time.Second
	s := New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		panic(err)
	}

	if err := s.Release(); err != nil {
		panic(err)
	}

}

func Demo2() {
	tickets := 0
	timeout := time.Duration(0)
	s := New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		if err != ErrNoTickets {
			panic(err)
		}

		// No tickets left, can't work :(
		os.Exit(2)
	}
}
