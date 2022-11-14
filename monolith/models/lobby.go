package models

type Lobby struct {
	lobby map[string][]chan *Game
}

func NewLobby() *Lobby {
	return &Lobby{
		lobby: make(map[string][]chan *Game),
	}
}

func (l *Lobby) Join(gameID string) chan *Game {
	if _, ok := l.lobby[gameID]; !ok {
		l.lobby[gameID] = make([]chan *Game, 0, 2)
	}

	ch := make(chan *Game)
	l.lobby[gameID] = append(l.lobby[gameID], ch)

	return ch
}

func (l *Lobby) Update(g *Game) {
	for _, lobbyChannel := range l.lobby[g.ID.String()] {
		lobbyChannel <- g
	}
}
