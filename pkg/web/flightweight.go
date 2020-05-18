package web

import "sync"

type Flyweight struct {
	sync.RWMutex
	data map[string]string
}

func NewFlyweight() Flyweight {
	return Flyweight{
		data: make(map[string]string),
	}
}

func (r *Flyweight) GetPage(tile string) string {
	r.RLock()
	page, ok := r.data[tile]
	r.RUnlock()
	if ok {
		return page
	}
	r.Lock()
	r.data[tile] = GetPage(tile)
	r.Unlock()
	return r.data[tile]
}
