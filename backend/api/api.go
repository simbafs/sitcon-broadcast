package api

import (
	"backend/internal/logger"
	"backend/internal/token"
	"backend/models/counter"
	"backend/sse"

	"github.com/danielgtaylor/huma/v2"
)

var log = logger.New("api")

func summary(summary string, description string, tags ...string) func(op *huma.Operation) {
	return func(op *huma.Operation) {
		op.Summary = summary
		op.Description = description
		op.Tags = tags
	}
}

func Route(api huma.API, h *Handler) {
	// counter
	huma.Get(api, "/api/counter/",
		h.GetAllCounter,
		summary("Get All Counter", "Get all counters.", "counter"),
	)
	huma.Get(api, "/api/counter/{name}",
		h.GetCounter,
		summary("Get Counter", "Get the counter by name.", "counter"),
	)
	huma.Put(api, "/api/counter/{name}",
		h.SetCounterInit,
		summary("Set Init Value", "Set the initial value and reset the counter.", "counter"),
		h.token.AuthHuma(api),
	)
	huma.Put(api, "/api/counter/{name}/start",
		h.CounterStart,
		summary("Start Counter", "Start the counter. It will reset the counter depend on the state.", "counter"),
		h.token.AuthHuma(api),
	)
	huma.Put(api, "/api/counter/{name}/stop",
		h.CounterStop,
		summary("Stop Counter", "Stop the counter. It will reset the counter when start it again.", "counter"),
		h.token.AuthHuma(api),
	)
	huma.Put(api, "/api/counter/{name}/reset",
		h.CounterReset,
		summary("Reset Counter", "Reset the counter", "counter"),
		h.token.AuthHuma(api),
	)

	// event
	huma.Get(api, "/api/event/",
		h.GetAllEvent,
		summary("Get All Events", "Get all events", "event"),
	)
	huma.Get(api, "/api/event/{name}",
		h.GetEvent,
		summary("Get Event", "Get event by name.", "event"),
	)
	huma.Get(api, "/api/event/{name}/session",
		h.GetEventSession,
		summary("Get Event Sessions", "Get sessions of event.", "event"),
	)
	huma.Post(api, "/api/event/",
		h.CreateEvent,
		summary("Create Event", "Create a new event.", "event"),
		h.token.AuthHuma(api),
	)
	huma.Put(api, "/api/event/{name}",
		h.UpdateEventScript,
		summary("Update Event Script", "Update event script by name", "event"),
		h.token.AuthHuma(api),
	)

	// now
	huma.Get(api, "/api/now/",
		h.GetNow,
		summary("Get Current Time", "Get the current time in unix timestamp in seconds", "now"),
	)
	huma.Post(api, "/api/now/",
		h.SetNow,
		summary("Set Current Time", "Set the current time in unix timestamp in seconds", "now"),
		h.token.AuthHuma(api),
	)
	huma.Delete(api, "/api/now/",
		h.ResetNow,
		summary("Reset Current Time", "Reset the current time to the actual current time.", "now"),
		h.token.AuthHuma(api),
	)

	// session
	huma.Get(api, "/api/session/{room}",
		h.GetAllSession,
		summary("Get Current Session in Room", "Get current session in room", "session"),
	)
	huma.Get(api, "/api/session/{room}/all",
		h.GetSessionInRoom,
		summary("Get All Sessions in Room", "Get all sessions in a room", "session"),
	)
	huma.Get(api, "/api/session/{room}/{id}",
		h.GetSessionByID,
		summary("Get Session by ID in Room", "Get a session by its ID in a room", "session"),
	)
	huma.Post(api, "/api/session/{room}/{id}",
		h.NextSession,
		summary("Set End Time of Session", "Set the end time of the current session and start time of the next session.", "session"),
		h.token.AuthHuma(api),
	)
	huma.Put(api, "/api/session/",
		h.SetAllSession,
		summary("Set Sessions", "Clear all sessions and set new sessions in database. Note that this API will not check if the session is valid.", "session"),
		h.token.AuthHuma(api),
	)
}

type Output[T any] struct {
	Body T `doc:"response body"`
}

type Handler struct {
	send     sse.Send
	token    *token.Token
	counters counter.CounterGroup
}

func NewHandler(send sse.Send, token *token.Token) *Handler {
	callback := func(name string) counter.Callback {
		return func(c *counter.Counter) {
			send <- sse.Msg{
				Topic: []string{"counter/" + name},
				Data:  *c,
			}
			log.Println(name, c)
		}
	}

	counters := counter.NewGroup([]string{
		"R0",
		"R1",
		"R2",
		"R3",
		"S",
	}, []counter.Callback{callback("R0"), callback("R1"), callback("R2"), callback("R3"), callback("S")})

	return &Handler{
		send:     send,
		token:    token,
		counters: counters,
	}
}
