package FinanransdonScriptLib

type Token struct {
	clip      string
	tokenType int
}

const (
	TokenTypeHexLiteral    = 0
	TokenTypeComment       = 1
	TokenTypeOpRepeat      = 2
	TokenTypeNumberLiteral = 3
	TokenLeftBracket       = 4
	TokenRightBracket      = 5
)
