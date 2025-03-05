package ticker

import (
	"context"
	"time"

	"backend/logger"
	"backend/middleware"
	"backend/models/now"
	"backend/models/room"
	"backend/models/session"
	"backend/models/special"
)

var log = logger.New("ticker")

type (
	Msg          any
	MsgPing      struct{}
	MsgNow       struct{}
	MsgCountdown string
	MsgOneCard   string
	MsgCard      struct {
		Room string
		ID   string
	}
	MsgSpecial string
)

func Listen(broadcast chan middleware.SSEMsg, quit chan struct{}, update chan Msg) {
	log.Println("Listening for updates")

	perSecond := time.NewTicker(1 * time.Second)
	perMinute := time.NewTicker(1 * time.Minute)

	for {
		select {
		case <-perSecond.C:
			UpdateCountdown(broadcast)
			UpdatePing(broadcast)
		case <-perMinute.C:
			UpdateNow(broadcast)
			UpdateAllCurrentCard(broadcast)
		case msg := <-update:
			switch msg := msg.(type) {
			case MsgPing:
				UpdatePing(broadcast)
			case MsgNow:
				UpdateNow(broadcast)
				UpdateAllCurrentCard(broadcast)
			case MsgCountdown:
				UpdateCountdown(broadcast)
			case MsgCard:
				UpdateCardInRoom(broadcast, msg.Room, msg.ID)
			case MsgSpecial:
				UpdateSepcial(broadcast, string(msg))
			}
		case <-quit:
			perSecond.Stop()
			perMinute.Stop()
			return
		}
	}
}

func UpdatePing(broadcast chan middleware.SSEMsg) {
	broadcast <- middleware.SSEMsg{
		Name: "ping",
		Data: time.Now().Unix(),
	}
}

func UpdateNow(broadcast chan middleware.SSEMsg) {
	broadcast <- middleware.SSEMsg{
		Name: "now",
		Data: now.Read(),
	}
}

// TODO: only send updates when changing.
func UpdateCountdown(broadcast chan middleware.SSEMsg) {
	for i, r := range room.Rooms {
		if r.State == room.COUNTING {
			r.Time -= 1

			if r.Time <= 0 {
				r.State = room.PAUSE
				r.Time = 0
			}

			room.Rooms[i] = r
		}

		broadcast <- middleware.SSEMsg{
			Name: "countdown-" + r.Name,
			Data: r,
		}
	}
}

var rooms = []string{"R0", "R1", "R2", "R3", "S"}

func UpdateAllCurrentCard(broadcast chan middleware.SSEMsg) {
	for _, room := range rooms {
		curr, err := session.ReadCurrentByRoom(context.Background(), room)
		if err != nil {
			log.Printf("failed to get current session of rooom %s: %s", room, err)
		}
		broadcast <- middleware.SSEMsg{
			Name: "card-current-" + room,
			Data: curr,
		}
	}
}

func UpdateOneCard(broadcast chan middleware.SSEMsg, id string) {
	curr, err := session.ReadByID(context.Background(), id)
	if err != nil {
		log.Printf("failed to get current session of rooom %s: %s", id, err)
	}

	broadcast <- middleware.SSEMsg{
		Name: "card-id-" + curr.ID,
		Data: curr,
	}
}

func UpdateCardInRoom(broadcast chan middleware.SSEMsg, room string, id string) {
	prev, target, next, err := session.ReadPrevNext(context.Background(), room, id)
	if err != nil {
		log.Println("failed to get prev, current, next session", err)
		return
	}

	curr, err := session.ReadCurrentByRoom(context.Background(), room)
	if err != nil {
		log.Println("failed to get current session", err)
		return
	}

	broadcast <- middleware.SSEMsg{
		Name: "card-current-" + room,
		Data: curr,
	}

	if target != nil {
		log.Println("curr", target)
		broadcast <- middleware.SSEMsg{
			Name: "card-" + room,
			Data: target,
		}
		broadcast <- middleware.SSEMsg{
			Name: "card-id-" + target.ID,
			Data: target,
		}
	}
	if prev != nil {
		log.Println("prev", prev)
		broadcast <- middleware.SSEMsg{
			Name: "card-" + room,
			Data: prev,
		}
		broadcast <- middleware.SSEMsg{
			Name: "card-id-" + prev.ID,
			Data: prev,
		}
	}
	if next != nil {
		log.Println("next", next)
		broadcast <- middleware.SSEMsg{
			Name: "card-" + room,
			Data: next,
		}
		broadcast <- middleware.SSEMsg{
			Name: "card-id-" + next.ID,
			Data: next,
		}
	}
}

func UpdateSepcial(broadcast chan middleware.SSEMsg, id string) {
	s, err := special.Read(context.Background(), id)
	if err != nil {
		log.Println("failed to get special", err)
		return
	}

	broadcast <- middleware.SSEMsg{
		Name: "special-" + id,
		Data: s.Data,
	}
}
