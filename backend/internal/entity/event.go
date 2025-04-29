package entity

type Event struct {
	name   string
	url    string
	script string
}

func NewEvent(name, url, script string) *Event {
	return &Event{
		name:   name,
		url:    url,
		script: script,
	}
}

func (e *Event) Name() string {
	return e.name
}

func (e *Event) URL() string {
	return e.url
}

func (e *Event) Script() string {
	return e.script
}

func (e *Event) SetURL(url string) {
	e.url = url
}

func (e *Event) SetScript(script string) {
	e.script = script
}
