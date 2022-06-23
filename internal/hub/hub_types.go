package hub

import "ws-noughts-and-crosses/pkg/vec"

type Direction struct {
	W bool `json:"w"`
	A bool `json:"a"`
	S bool `json:"s"`
	D bool `json:"d"`
}

type PlayerActionMessage struct {
	ID        string    `json:"id"`
	Direction Direction `json:"direction"`
	Pos       vec.Vec   `json:"pos"`
	Vel       vec.Vec   `json:"vel"`
}

type CurrentStateMessage struct {
	Update []PlayerActionMessage `json:"update"`
}
