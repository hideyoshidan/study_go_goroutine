package mutex

import "sync"

type WiMutex struct {
	sync.Mutex
	dictonary map[string]int
}

func NewWiMutex() *WiMutex {
	return &WiMutex{
		dictonary: make(map[string]int),
	}
}

func (m *WiMutex) Set(key string) {
	m.Lock()
	m.dictonary[key]++
	m.Unlock()
}

func (m *WiMutex) Get(key string) int {
	m.Lock()
	val := m.dictonary[key]
	m.Unlock()
	return val
}
