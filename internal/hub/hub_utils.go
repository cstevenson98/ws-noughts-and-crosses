package hub

import (
	"encoding/json"
	"math/rand"
	"time"
)

func (h *Hub) PlayerCount() (sum int) {
	for game := range h.Games {
		sum += len(game.Players)
	}
	return
}

func (h *Hub) LogOnInterval(interval time.Duration) {
	for range time.Tick(interval) {
		for game := range h.Games {
			for player := range game.Players {
				randomMessage := CurrentStateMessage{
					Positions: [][2]float64{{rand.Float64(), rand.Float64()}},
				}
				payload, _ := json.Marshal(randomMessage)
				player.Stream <- payload
			}
			game.t0 += interval.Seconds()
		}
	}
}