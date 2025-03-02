package ticker

import (
	"context"
	"time"

	"backend/logger"
	"backend/middleware"
	"backend/models/now"
	"backend/models/room"
	"backend/models/session"
)

var log = logger.New("ticker")

type Msg int

const (
	MsgPing Msg = 0x01 << iota
	MsgNow
	MsgCountdown
	MsgCard
	MsgAll = MsgPing | MsgNow | MsgCountdown | MsgCard
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
			UpdateCard(broadcast)
		case msg := <-update:
			if msg&MsgPing == MsgPing {
				UpdatePing(broadcast)
			}
			if msg&MsgNow == MsgNow {
				UpdateNow(broadcast)
			}
			if msg&MsgCountdown == MsgCountdown {
				UpdateCountdown(broadcast)
			}
			if msg&MsgCard == MsgCard {
				UpdateCard(broadcast)
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

func UpdateCard(broadcast chan middleware.SSEMsg) {
	for _, room := range rooms {
		if now, err := session.ReadCurrentByRoom(context.Background(), room); err != nil {
			broadcast <- middleware.SSEMsg{
				Name: "card-" + room,
				Data: now,
			}
		}
	}
}
