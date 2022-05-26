package hub

type CurrentStateMessage struct {
	MyPosition [2]float64   `json:"my_position"`
	Positions  [][2]float64 `json:"positions"`
}

// UserInputEventMessage tells the direction the player is accelerating
type UserInputEventMessage struct {
	W bool `json:"w"`
	A bool `json:"a"`
	S bool `json:"s"`
	D bool `json:"d"`
}