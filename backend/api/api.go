package api

import (
	"backend/api/card"
	"backend/api/now"
	aRoom "backend/api/room"

	"github.com/gin-gonic/gin"
)

// func timer(quit chan struct{}, io websocket.IO) {
// 	tricker := time.NewTicker(1 * time.Second)
// 	for {
// 		select {
// 		case <-tricker.C:
// 			for i, r := range room.Rooms {
// 				if r.State == room.PAUSE {
// 					continue
// 				}
//
// 				r.Time -= 1
// 				if r.Time <= 0 {
// 					r.State = room.PAUSE
// 					r.Time = 0
// 				}
//
// 				room.Rooms[i] = r
// 			}
// 			// log.Printf("%#v\n", rooms )
// 			// data, err := json.Marshal(rooms)
// 			data, err := json.Marshal(gin.H{
// 				"rooms":      room.Rooms,
// 				"serverTime": time.Now(),
// 			})
// 			if err != nil {
// 				log.Println(err)
// 				continue
// 			}
//
// 			io.Broadcast(data)
// 		case <-quit:
// 			tricker.Stop()
// 			return
// 		}
// 	}
// }

func Route(r *gin.Engine) {
	api := r.Group("/api")

	card.Route(api)
	now.Route(api)
	aRoom.Route(api)

	// TODO:
	// go timer(quit, io)
}
