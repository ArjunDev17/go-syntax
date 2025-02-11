package services

import "tic-tac-toe/core"

type GameService struct {
	game *core.Game
}

func NewGameService(size int, players []core.Player) *GameService {
	return &GameService{
		game: core.NewGame(size, players),
	}
}

func (gs *GameService) Start() {
	gs.game.Play()
}
