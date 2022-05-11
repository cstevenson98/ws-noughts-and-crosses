package hub

type CurrentStateMessage struct {
	Positions [][2]float64 `json:"positions"`
}