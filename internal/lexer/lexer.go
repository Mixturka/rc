package lexer

import (
	"slices"
	"unicode"

	"github.com/Mixturka/rc/internal/lexer/token"
	"github.com/Mixturka/rc/internal/pkg/scope"
)

type LexerError int

const (
	CommentSkipped LexerError = iota
	NewLineSkipped
	UnexpectedSymbol
	WhitespaceSkipped
	TabSkipped
)

func (le LexerError) Error() string {
	return "" // TODO: add proper error output
}

type Lexer struct {
	src  []rune
	pos  int
	line int
}

func NewLexer(src []rune) *Lexer {
	return &Lexer{
		src: src,
	}
}

func (l *Lexer) Tokenize() (tokens []token.Token, err error) {
	for l.pos < len(l.src) {
		if tok, err := l.scanToken(); err != nil {
			if err == CommentSkipped || err == NewLineSkipped || err == WhitespaceSkipped || err == TabSkipped {
				continue
			}
			return nil, err
		} else {
			tokens = append(tokens, tok)
		}
	}

	return tokens, nil
}

func (l *Lexer) scanToken() (token.Token, error) {
	ch := l.src[l.pos]
	scope := scope.Scope{Start: l.pos, End: l.pos}
	l.pos++

	switch ch {
	case '(':
		return token.Token{Type: token.LeftParen, Scope: scope}, nil
	case ')':
		return token.Token{Type: token.RightParen, Scope: scope}, nil
	case '{':
		return token.Token{Type: token.LeftBrace, Scope: scope}, nil
	case '}':
		return token.Token{Type: token.RightBrace, Scope: scope}, nil
	case ':':
		return token.Token{Type: token.Colon, Scope: scope}, nil
	case ';':
		return token.Token{Type: token.Semicolon, Scope: scope}, nil
	case '*':
		if tok, ok := l.expectNext(ExpectedInfo{'=', token.StarAssign}); ok {
			return tok, nil
		}
		return token.Token{Type: token.Star, Scope: scope}, nil
	case '-':
		if tok, ok := l.expectNext(ExpectedInfo{'>', token.Arrow}, ExpectedInfo{'=', token.MinusAssign}, ExpectedInfo{'-', token.MinusMinus}); ok {
			return tok, nil
		} else if l.pos < len(l.src) && unicode.IsDigit(l.src[l.pos+1]) {
			return l.scanNumber(scope.Start)
		}

		return token.Token{Type: token.Minus, Scope: scope}, nil
	case '+':
		if tok, ok := l.expectNext(ExpectedInfo{'=', token.PlusAssign}, ExpectedInfo{'+', token.PlusPlus}); ok {
			return tok, nil
		}
		return token.Token{Type: token.Plus, Scope: scope}, nil
	case '/':
		if tok, ok := l.expectNext(ExpectedInfo{'=', token.SlashAssign}); ok {
			return tok, nil
		} else if l.pos < len(l.src) {
			switch l.src[l.pos] {
			case '/':
				l.pos++
				l.skipOneLineComment()
				return token.Token{}, CommentSkipped
			case '*':
				l.pos++
				l.skipMultiLineComment()
				return token.Token{}, CommentSkipped
			default:
				return token.Token{Type: token.Slash, Scope: scope}, nil
			}
		}
		return token.Token{Type: token.Slash, Scope: scope}, nil
	case '=':
		if tok, ok := l.expectNext(ExpectedInfo{'=', token.Equals}); ok {
			return tok, nil
		}
		return token.Token{Type: token.Assign, Scope: scope}, nil
	case '!':
		if tok, ok := l.expectNext(ExpectedInfo{'=', token.NotEquals}); ok {
			return tok, nil
		}
		return token.Token{Type: token.Not, Scope: scope}, nil
	case '\n':
		l.line++
		return token.Token{}, NewLineSkipped
	case ' ':
		return token.Token{}, WhitespaceSkipped
	case '\t':
		return token.Token{}, TabSkipped
	default:
		switch {
		case unicode.IsLetter(ch) || ch == '_':
			tok, err := l.scanIdentifier(scope.Start)
			if err != nil {
				return token.Token{}, nil
			}

			if keyword, ok := l.checkKeyword(tok); ok {
				return keyword, nil
			}
			return tok, nil
		case unicode.IsDigit(ch):
			tok, err := l.scanNumber(scope.Start)
			if err != nil {
				return token.Token{}, nil
			}
			return tok, nil
		default:
			return token.Token{}, UnexpectedSymbol
		}
	}
}

type ExpectedInfo struct {
	Ch      rune
	TokType token.TokenType // type to return in case of success
}

func (l *Lexer) expectNext(options ...ExpectedInfo) (token.Token, bool) {
	if l.pos >= len(l.src) {
		return token.Token{}, false
	}

	next := l.src[l.pos]
	for _, v := range options {
		if next == v.Ch {
			l.pos++
			return token.Token{Type: v.TokType, Scope: scope.Scope{Start: l.pos - 2, End: l.pos - 1}}, true
		}
	}

	return token.Token{}, false
}

func (l *Lexer) skipOneLineComment() {
	for l.src[l.pos] != '\n' {
		l.pos++
	}
	l.pos++
}

func (l *Lexer) skipMultiLineComment() {
	for l.pos < len(l.src) {
		if l.pos < len(l.src)-1 && slices.Equal(l.src[l.pos:l.pos+2], []rune("*/")) {
			l.pos += 2
			return
		}
		if l.src[l.pos] == '\n' {
			l.line++
		}
		l.pos++
	}
}

func (l *Lexer) scanIdentifier(scopeStart int) (token.Token, error) {
	if l.pos >= len(l.src) {
		return token.Token{Type: token.Identifier, Scope: scope.Scope{Start: scopeStart, End: scopeStart}}, nil
	}
	ch := l.src[l.pos]

	for l.pos < len(l.src)-1 && (unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_') {
		l.pos++
		ch = l.src[l.pos]
	}
	if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' {
		l.pos++
	}

	return token.Token{Type: token.Identifier, Scope: scope.Scope{Start: scopeStart, End: l.pos - 1}}, nil
}

func (l *Lexer) scanNumber(scopeStart int) (token.Token, error) {
	if l.pos >= len(l.src) {
		return token.Token{Type: token.Identifier, Scope: scope.Scope{Start: scopeStart, End: scopeStart}}, nil
	}
	ch := l.src[l.pos]

	for l.pos < len(l.src)-1 && unicode.IsDigit(ch) {
		l.pos++
		ch = l.src[l.pos]
	}
	if unicode.IsDigit(ch) {
		l.pos++
	}

	return token.Token{Type: token.IntegerNumber, Scope: scope.Scope{Start: scopeStart, End: l.pos - 1}}, nil
}

func (l *Lexer) checkKeyword(tok token.Token) (token.Token, bool) {
	tokStr := string(l.src[tok.Scope.Start : tok.Scope.End+1])

	switch tokStr {
	case "fn":
		return token.Token{Type: token.Fn, Scope: tok.Scope}, true
	case "return":
		return token.Token{Type: token.Return, Scope: tok.Scope}, true
	default:
		return token.Token{}, false
	}
}
