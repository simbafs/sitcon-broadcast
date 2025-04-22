package counter

type CounterGroup map[string]*Counter

type Callback func(counter *Counter)

func NewGroup(names []string, callbacks []Callback) CounterGroup {
	if len(names) != len(callbacks) {
		panic("length of names and callbacks must be equal")
	}
	cg := make(CounterGroup, len(names))
	for i, name := range names {
		cg[name] = NewCounter(60, callbacks[i])
	}
	return cg
}

func (cg CounterGroup) Get(name string) *Counter {
	return cg[name]
}
