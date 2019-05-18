package syncevent

import (
	"sync"
)

var _ SyncEvent = &syncEvent{}

// NewSyncEvent creates a new sync event flag
func NewSyncEvent(initValue bool) SyncEvent {
	mutex := &sync.RWMutex{}

	return &syncEvent{
		flagMutex: mutex,
		flag:      initValue,
		cond:      sync.NewCond(mutex.RLocker()),
	}
}

type syncEvent struct {
	flagMutex *sync.RWMutex
	flag      bool
	cond      *sync.Cond
}

func (s *syncEvent) IsSet() bool {
	s.flagMutex.RLock()
	defer s.flagMutex.RUnlock()
	return s.flag
}

func (s *syncEvent) Set() {
	s.flagMutex.Lock()
	defer s.flagMutex.Unlock()
	s.flag = true
	s.cond.Broadcast()
}

func (s *syncEvent) Wait() {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	for !s.flag {
		s.cond.Wait()
	}
}

func (s *syncEvent) Reset() {
	s.flagMutex.Lock()
	defer s.flagMutex.Unlock()
	s.flag = false
}
