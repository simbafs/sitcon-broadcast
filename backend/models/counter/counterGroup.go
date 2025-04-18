package counter

import "backend/internal/logger"

type CounterGroup []*Counter

var log = logger.New("counter")

func NewGroup(names []string) CounterGroup {
	cg := make(CounterGroup, len(names))
	for i, name := range names {
		cg[i] = &Counter{
			Name:   name,
			Status: StatusStopped,
			Init:   0,
			Count:  0,
		}
	}
	return cg
}

func (cg CounterGroup) Update() {
	for i := range cg {
		cg[i].Update()
	}
}

func (cg CounterGroup) Get(name string) *Counter {
	for _, c := range cg {
		log.Println(c.Name, name)
		if c.Name == name {
			return c
		}
	}
	return nil
}
