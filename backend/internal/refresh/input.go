package refresh

type Input struct {
	Sessions     []Session     `json:"sessions"`
	Speakers     []Speaker     `json:"speakers"`
	SessionTypes []SessionType `json:"session_types"`
	Rooms        []Room        `json:"rooms"`
	Tags         []Tag         `json:"tags"`
}

type Session struct {
	ID        string   `json:"id"`
	Type      string   `json:"type"`
	Room      string   `json:"room"`
	Broadcast []string `json:"broadcast"`
	Start     string   `json:"start"`
	End       string   `json:"end"`
	QA        string   `json:"qa,omitempty"`
	Slide     string   `json:"slide,omitempty"`
	CoWrite   string   `json:"co_write,omitempty"`
	Record    string   `json:"record"`
	Live      string   `json:"live"`
	URI       string   `json:"uri,omitempty"`
	Zh        struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"zh"`
	Speakers []string `json:"speakers"`
	Tags     []string `json:"tags"`
}

type Speaker struct {
	ID string `json:"id"`
	Zh struct {
		Name string `json:"name"`
	} `json:"zh"`
}

type Room struct {
	ID string `json:"id"`
	Zh struct {
		Name string `json:"name"`
	} `json:"zh"`
}

type SessionType struct {
	ID string `json:"id"`
	Zh struct {
		Name string `json:"name"`
	} `json:"zh"`
}

type Tag struct {
	ID string `json:"id"`
	Zh struct {
		Name string `json:"name"`
	} `json:"zh"`
}
