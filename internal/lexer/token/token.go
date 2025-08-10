package token

import "github.com/Mixturka/rc/internal/pkg/scope"

type TokenType int

const (
	LeftParen   TokenType = iota // (
	RightParen                   // )
	LeftBrace                    // {
	RightBrace                   // }
	Arrow                        // ->
	Colon                        // :
	Semicolon                    // ;
	Comma                        // ,
	Star                         // *
	Minus                        // -
	Plus                         // +
	Slash                        // /
	MinusAssign                  // -=
	MinusMinus                   // --
	PlusAssign                   // +=
	PlusPlus                     // ++
	StarAssign                   // *=
	SlashAssign                  // /=
	Assign                       // =
	Equals                       // ==
	NotEquals                    // !=
	Not                          // !
	Tilde                        // ~
	Identifier
	IntegerNumber
	Fn
	Return
	Eof
)

type Token struct {
	Type  TokenType
	Scope scope.Scope
}

func (tt TokenType) BindingPower() int {
	switch tt {
	case Plus:
		fallthrough
	case Minus:
		return 1
	case Slash:
		fallthrough
	case Star:
		return 2
	default:
		return 0
	}
}
