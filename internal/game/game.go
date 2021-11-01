package game

type Turn struct {
	encodedTurn string
}

type IState interface {
	// ?
}

type IPlayer interface {
	ReadPump()
	WritePump()
	GetGame() *IGame
	MessagePlayer(message []byte)
}

type IGame interface {
	ProcessTurn(turn Turn)
	BroadcastState()
}
