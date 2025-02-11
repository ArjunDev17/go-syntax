package core

type Player struct {
	Name   string
	Symbol PlayerSymbol
}

func NewPlayer(name string, symbol PlayerSymbol) Player {
	return Player{
		Name:   name,
		Symbol: symbol,
	}
}
