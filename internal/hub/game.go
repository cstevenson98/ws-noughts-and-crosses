package hub

const (
	GameWaiting = "Waiting"
)

type Turn struct {
	player      IPlayer
	encodedTurn []byte
}

type IState interface {
	// ?
}

type IPlayer interface {
	ReadPump()
	WritePump()
}

type IGame interface {
}