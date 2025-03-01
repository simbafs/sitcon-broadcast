package ticker

import (
	"time"

	"backend/logger"
	"backend/middleware"
	"backend/models/now"
	"backend/models/room"
	"backend/models/session"
)

var log = logger.New("ticker")

func Listen(broadcast chan middleware.SSEMsg, quit chan struct{}, updateAll chan struct{}) {
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
		case <-updateAll:
			UpdateCountdown(broadcast)
			UpdatePing(broadcast)
			UpdateNow(broadcast)
			UpdateCard(broadcast)
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
		Data: now.GetNow(),
	}
}

func UpdateCountdown(broadcast chan middleware.SSEMsg) {
	for i, r := range room.Rooms {
		if r.State == room.PAUSE {
			continue
		}

		r.Time -= 1

		if r.Time <= 0 {
			r.State = room.PAUSE
			r.Time = 0
		}

		room.Rooms[i] = r

		broadcast <- middleware.SSEMsg{
			Name: "countdown-" + r.Name,
			Data: r,
		}
	}
}

func UpdateCard(broadcast chan middleware.SSEMsg) {
	for name, r := range session.Data.Rooms {
		if now, ok := r.GetNow(); ok {
			broadcast <- middleware.SSEMsg{
				Name: "card-" + name,
				Data: now,
			}
		}
	}
}
