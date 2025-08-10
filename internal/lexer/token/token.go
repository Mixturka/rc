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
	Identifier
	IntegerNumber
	Fn
	Return
)

type Token struct {
	Type  TokenType
	Scope scope.Scope
}
