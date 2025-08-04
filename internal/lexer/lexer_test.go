package lexer_test

import (
	"slices"
	"testing"

	"github.com/Mixturka/rc/internal/lexer"
	"github.com/Mixturka/rc/internal/lexer/token"
	"github.com/Mixturka/rc/internal/pkg/scope"
)

func TestLexLeftParen(t *testing.T) {
	l := lexer.NewLexer([]rune{'('})
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.LeftParen, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v got %v", correctTokenSlice, toks)
	}
}

func TestLexRightParen(t *testing.T) {
	l := lexer.NewLexer([]rune{')'})
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.RightParen, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v got %v", correctTokenSlice, toks)
	}
}

func TestLexLeftBrace(t *testing.T) {
	l := lexer.NewLexer([]rune{'{'})
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.LeftBrace, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexRightBrace(t *testing.T) {
	l := lexer.NewLexer([]rune{'}'})
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.RightBrace, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexArrow(t *testing.T) {
	l := lexer.NewLexer([]rune("->"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Arrow, Scope: scope.Scope{Start: 0, End: 1}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexColon(t *testing.T) {
	l := lexer.NewLexer([]rune(":"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Colon, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexSemicolon(t *testing.T) {
	l := lexer.NewLexer([]rune(";"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Semicolon, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexStar(t *testing.T) {
	l := lexer.NewLexer([]rune("*"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Star, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexMinus(t *testing.T) {
	l := lexer.NewLexer([]rune("-"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Minus, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexPlus(t *testing.T) {
	l := lexer.NewLexer([]rune("+"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Plus, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexSlash(t *testing.T) {
	l := lexer.NewLexer([]rune("/"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Slash, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexSkipOneLineComment(t *testing.T) {
	l := lexer.NewLexer([]rune("// blah blah blah \n ("))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.LeftParen, Scope: scope.Scope{Start: 20, End: 20}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexSkipMulilineComment(t *testing.T) {
	l := lexer.NewLexer([]rune("/* blah blah blah \n super blah */("))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.LeftParen, Scope: scope.Scope{Start: 33, End: 33}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexMinusAssign(t *testing.T) {
	l := lexer.NewLexer([]rune("-="))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.MinusAssign, Scope: scope.Scope{Start: 0, End: 1}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexMinusMinus(t *testing.T) {
	l := lexer.NewLexer([]rune("--"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.MinusMinus, Scope: scope.Scope{Start: 0, End: 1}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexPlusAssign(t *testing.T) {
	l := lexer.NewLexer([]rune("+="))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.PlusAssign, Scope: scope.Scope{Start: 0, End: 1}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexPlusPlus(t *testing.T) {
	l := lexer.NewLexer([]rune("++"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.PlusPlus, Scope: scope.Scope{Start: 0, End: 1}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexStarAssign(t *testing.T) {
	l := lexer.NewLexer([]rune("*="))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.StarAssign, Scope: scope.Scope{Start: 0, End: 1}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexSlashAssign(t *testing.T) {
	l := lexer.NewLexer([]rune("/="))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.SlashAssign, Scope: scope.Scope{Start: 0, End: 1}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexAssign(t *testing.T) {
	l := lexer.NewLexer([]rune("="))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Assign, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexEquals(t *testing.T) {
	l := lexer.NewLexer([]rune("=="))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Equals, Scope: scope.Scope{Start: 0, End: 1}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexNotEquals(t *testing.T) {
	l := lexer.NewLexer([]rune("!="))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.NotEquals, Scope: scope.Scope{Start: 0, End: 1}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexNot(t *testing.T) {
	l := lexer.NewLexer([]rune("!"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Not, Scope: scope.Scope{Start: 0, End: 0}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexIdentifier1(t *testing.T) {
	l := lexer.NewLexer([]rune("identifier"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Identifier, Scope: scope.Scope{Start: 0, End: 9}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexIdentifier2(t *testing.T) {
	l := lexer.NewLexer([]rune("_identifier_"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Identifier, Scope: scope.Scope{Start: 0, End: 11}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexIdentifier3(t *testing.T) {
	l := lexer.NewLexer([]rune("identifier_123_cool"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Identifier, Scope: scope.Scope{Start: 0, End: 18}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexDecimalPositiveIntegerNumber(t *testing.T) {
	l := lexer.NewLexer([]rune("123"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.IntegerNumber, Scope: scope.Scope{Start: 0, End: 2}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexDecimalNegativeIntegerNumber(t *testing.T) {
	l := lexer.NewLexer([]rune("-123"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.IntegerNumber, Scope: scope.Scope{Start: 0, End: 3}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexFnKeyword(t *testing.T) {
	l := lexer.NewLexer([]rune("fn"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Fn, Scope: scope.Scope{Start: 0, End: 1}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexReturnKeyword(t *testing.T) {
	l := lexer.NewLexer([]rune("return"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Return, Scope: scope.Scope{Start: 0, End: 5}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}

func TestLexMainFunctionWithOneReturnInteger(t *testing.T) {
	l := lexer.NewLexer([]rune("fn main() -> i32 {\n\treturn 23;\n}"))
	toks, err := l.Tokenize()
	correctTokenSlice := []token.Token{
		{Type: token.Fn, Scope: scope.Scope{Start: 0, End: 1}}, {Type: token.Identifier, Scope: scope.Scope{Start: 3, End: 6}},
		{Type: token.LeftParen, Scope: scope.Scope{Start: 7, End: 7}}, {Type: token.RightParen, Scope: scope.Scope{Start: 8, End: 8}},
		{Type: token.Arrow, Scope: scope.Scope{Start: 10, End: 11}}, {Type: token.Identifier, Scope: scope.Scope{Start: 13, End: 15}},
		{Type: token.LeftBrace, Scope: scope.Scope{Start: 17, End: 17}}, {Type: token.Return, Scope: scope.Scope{Start: 20, End: 25}},
		{Type: token.IntegerNumber, Scope: scope.Scope{Start: 27, End: 28}}, {Type: token.Semicolon, Scope: scope.Scope{Start: 29, End: 29}},
		{Type: token.RightBrace, Scope: scope.Scope{Start: 31, End: 31}},
	}
	if !slices.Equal(toks, correctTokenSlice) || err != nil {
		t.Errorf("Expected: %v, got %v", correctTokenSlice, toks)
	}
}
