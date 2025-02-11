package adapters

import (
	"tic-tac-toe/core"
	"tic-tac-toe/services"
)

func StartCLI() {
	players := []core.Player{
		core.NewPlayer("Arjun", &core.SymbolX{}),
		core.NewPlayer("Singh", &core.SymbolO{}),
	}

	gameService := services.NewGameService(3, players)
	gameService.Start()
}
