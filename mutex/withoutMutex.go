package mutex

type WioutMutex struct {
	dictonary map[string]int
}

func NewWioutMutex() *WioutMutex {
	return &WioutMutex{
		dictonary: make(map[string]int),
	}
}

func (m *WioutMutex) Set(key string) {
	m.dictonary[key]++
}

func (m *WioutMutex) Get(key string) int {
	return m.dictonary[key]
}
