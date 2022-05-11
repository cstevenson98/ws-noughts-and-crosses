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
				payload, _ := json.Marshal([2]float64{rand.Float64() * 100, rand.Float64() * 100})
				player.Stream <- payload
			}
			game.t0 += interval.Seconds()
		}
	}
}