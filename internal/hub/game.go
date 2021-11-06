package hub

const (
	GameWaiting = "Waiting"
	GamePlayer1 = "Player 1's turn!"
	GamePlayer2 = "Player 2's turn!"
	GameOver    = "Game Over!"
)

type Turn struct {
	player IPlayer
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