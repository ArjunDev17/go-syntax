package core

type PlayerSymbol interface {
	GetSymbol() string
}

type SymbolType string

const (
	X SymbolType = "X"
	O SymbolType = "O"
)

type SymbolX struct{}

func (s *SymbolX) GetSymbol() string {
	return string(X)
}

type SymbolO struct{}

func (s *SymbolO) GetSymbol() string {
	return string(O)
}
