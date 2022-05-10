package hub

const MaxPlayers = 20

type Game struct {
	IGame
	Players []*Player
	Status  string
	t0      float64
}

func (g *Game) AddClient(player *Player) error {
	g.Players = append(g.Players, player)
	player.Game = g
	return nil
}

func (g *Game) SlotsFree() int {
	return MaxPlayers - len(g.Players)
}
