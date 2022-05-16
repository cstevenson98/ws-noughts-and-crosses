package hub

type CurrentStateMessage struct {
	MyPosition [2]float64   `json:"my_position"`
	Positions  [][2]float64 `json:"positions"`
}
