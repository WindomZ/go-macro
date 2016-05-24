package store

import (
	"container/list"
	"sync"
	"time"
)

type Store interface {
	Set(id string, digits []byte)
	Get(id string, clear bool) (digits []byte)
	Remove(id string)
}

type idByTimeValue struct {
	timestamp time.Time
	id        string
}

type memoryStore struct {
	sync.RWMutex
	digitsById map[string][]byte
	idByTime   *list.List
	numStored  int
	collectNum int
	expiration time.Duration
}

func NewMemoryStore(collectNum int, expiration time.Duration) Store {
	s := new(memoryStore)
	s.digitsById = make(map[string][]byte)
	s.idByTime = list.New()
	s.collectNum = collectNum
	s.expiration = expiration
	return s
}

func (s *memoryStore) Set(id string, digits []byte) {
	s.Lock()
	s.digitsById[id] = digits
	s.idByTime.PushBack(idByTimeValue{time.Now(), id})
	s.numStored++
	if s.numStored <= s.collectNum {
		s.Unlock()
		return
	}
	s.Unlock()
	go s.collect()
}

func (s *memoryStore) Get(id string, clear bool) (digits []byte) {
	if !clear {
		s.RLock()
		defer s.RUnlock()
	} else {
		s.Lock()
		defer s.Unlock()
	}
	digits, ok := s.digitsById[id]
	if !ok {
		return
	} else if clear {
		delete(s.digitsById, id)
	}
	return
}

func (s *memoryStore) collect() {
	now := time.Now()
	s.Lock()
	defer s.Unlock()
	s.numStored = 0
	for e := s.idByTime.Front(); e != nil; {
		ev, ok := e.Value.(idByTimeValue)
		if !ok {
			return
		} else if ev.timestamp.Add(s.expiration).Before(now) {
			delete(s.digitsById, ev.id)
			next := e.Next()
			s.idByTime.Remove(e)
			e = next
		} else {
			return
		}
	}
}

func (s *memoryStore) Remove(id string) {
	if len(id) == 0 {
		return
	}
	s.Lock()
	defer s.Unlock()
	delete(s.digitsById, id)
}
