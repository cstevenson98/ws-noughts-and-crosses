package hub

import "ws-noughts-and-crosses/pkg/vec"

type CurrentStateMessage struct {
	MyPosition [2]float64 `json:"my_position"`
	Positions  []vec.Vec  `json:"positions"`
}

// UserInputEventMessage tells the direction the player is accelerating
type UserInputEventMessage struct {
	W         bool `json:"w"`
	A         bool `json:"a"`
	S         bool `json:"s"`
	D         bool `json:"d"`
	Timestamp float64
}
