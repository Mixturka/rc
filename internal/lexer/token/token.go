package token

import "github.com/Mixturka/rc/internal/pkg/scope"

type TokenType int

const (
	LeftParen          TokenType = iota // (
	RightParen                          // )
	LeftBrace                           // {
	RightBrace                          // }
	Arrow                               // ->
	Colon                               // :
	Semicolon                           // ;
	Comma                               // ,
	Star                                // *
	Minus                               // -
	Plus                                // +
	Slash                               // /
	MinusAssign                         // -=
	MinusMinus                          // --
	PlusAssign                          // +=
	PlusPlus                            // ++
	StarAssign                          // *=
	SlashAssign                         // /=
	Assign                              // =
	Equals                              // ==
	NotEquals                           // !=
	Not                                 // !
	Tilde                               // ~
	Percent                             // %
	Ampersand                           // &
	AmpersandAmpersand                  // &&
	Bar                                 // |
	BarBar                              // ||
	Less                                // <
	Greater                             // >
	LessEqual                           // <=
	GreaterEqual                        // >=
	Identifier
	IntegerNumber
	Fn
	Return
	Eof
)

func (tt TokenType) IsOp() bool {
	switch tt {
	case Plus:
		fallthrough
	case Star:
		fallthrough
	case Minus:
		fallthrough
	case Slash:
		fallthrough
	case MinusMinus:
		fallthrough
	case Tilde:
		fallthrough
	case LeftParen:
		fallthrough
	case RightParen:
		fallthrough
	case Percent:
		fallthrough
	case Ampersand:
		fallthrough
	case AmpersandAmpersand:
		fallthrough
	case Bar:
		fallthrough
	case BarBar:
		fallthrough
	case Less:
		fallthrough
	case LessEqual:
		fallthrough
	case Greater:
		fallthrough
	case GreaterEqual:
		fallthrough
	case PlusPlus:
		return true
	}

	return false
}

type Token struct {
	Type  TokenType
	Scope scope.Scope
}
