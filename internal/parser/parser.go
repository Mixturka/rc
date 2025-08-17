package parser

import (
	"fmt"
	"log"

	"github.com/Mixturka/rc/internal/erremitter"
	"github.com/Mixturka/rc/internal/lexer/token"
	"github.com/Mixturka/rc/internal/parser/ast"
)

var (
	funcSignSyncSet = map[token.TokenType]struct{}{
		token.LeftBrace:  {},
		token.RightBrace: {},
		token.Arrow:      {},
		token.Semicolon:  {},
		token.LeftParen:  {},
	}
	stmtSyncSet = map[token.TokenType]struct{}{
		token.Semicolon:  {},
		token.RightBrace: {},
	}
	exprSyncSet = map[token.TokenType]struct{}{
		token.Comma:      {},
		token.RightParen: {},
	}
)

type Parser struct {
	inErr            bool
	pos              int
	tokens           []token.Token
	currentSyncStack []map[token.TokenType]struct{}
	errEmitter       *erremitter.ErrEmitter
	src              []rune
}

func NewParser(tokens []token.Token, errEmitter *erremitter.ErrEmitter, src []rune) Parser {
	return Parser{
		tokens:           tokens,
		errEmitter:       errEmitter,
		currentSyncStack: make([]map[token.TokenType]struct{}, 0),
		src:              src,
	}
}

func (p *Parser) next() *token.Token {
	if p.pos >= len(p.tokens) {
		return nil
	}

	token := p.tokens[p.pos]
	p.pos++
	return &token
}

func (p *Parser) peek() *token.Token {
	if p.pos >= len(p.tokens) {
		return nil
	}

	return &p.tokens[p.pos]
}

func (p *Parser) Parse() *ast.Program {
	return p.parseProgram()
}

func (p *Parser) parseProgram() *ast.Program {
	fn := p.ParseFunction()
	return &ast.Program{Function: *fn}
}

func (p *Parser) ParseFunction() *ast.Func {
	p.pushSyncStack(funcSignSyncSet)
	defer p.popSyncStack()

	if _, ok := p.expectAndConsumeToken(token.Fn); !ok {
		log.Fatalf("expected 'fn' keyword")
	}
	var funcName token.Token
	var ok bool
	funcName, ok = p.expectAndConsumeToken(token.Identifier)
	if !ok {
		log.Fatalf("expected function name")
	}

	if _, ok := p.expectAndConsumeToken(token.LeftParen); !ok {
		log.Fatalf("expected '('")
	}
	if _, ok := p.expectAndConsumeToken(token.RightParen); !ok {
		log.Fatalf("expected ')'")
	}
	if _, ok := p.expectAndConsumeToken(token.Arrow); !ok {
		log.Fatalf("expected '->'")
	}
	if _, ok := p.expectAndConsumeToken(token.Identifier); !ok {
		log.Fatalf("expected return type")
	}
	if _, ok := p.expectAndConsumeToken(token.LeftBrace); !ok {
		log.Fatalf("expected '{' before function body")
	}

	stmt := p.parseStatement()

	if _, ok := p.expectAndConsumeToken(token.RightBrace); !ok {
		log.Fatalf("expected '}'")
	}

	return &ast.Func{
		Name: funcName,
		Body: stmt,
	}
}

func (p *Parser) parseStatement() ast.Stmt {
	p.pushSyncStack(stmtSyncSet)
	defer p.popSyncStack()

	if _, ok := p.expectAndConsumeToken(token.Return); !ok {
		log.Fatalf("expected 'return'")
	}

	expr := p.parseExpression(0)

	if _, ok := p.expectAndConsumeToken(token.Semicolon); !ok {
		log.Fatalf("expected ';' in the end of statement")
	}

	return &ast.ReturnStmt{Expr: expr}
}

// minBp - minimal BindingPower for Pratt's Parser loop
func (p *Parser) parseExpression(minBp uint8) ast.Expr {
	p.pushSyncStack(exprSyncSet)
	defer p.popSyncStack()

	tok := p.next()
	fmt.Printf("Cur expr token: %v\n", tok)
	var lhs ast.Expr = &ast.ConstExpr{Value: *tok}
	if tok.Type.IsOp() {
		if tok.Type == token.LeftParen {
			lhs = p.parseExpression(0)
			if p.next().Type != token.RightParen {
				log.Fatal("expected ')' after expression")
			}
		} else {
			_, rBp := prefixBindingPower(tok.Type)
			rhs := p.parseExpression(rBp)
			lhs = &ast.UnaryExpr{Op: *tok, Rhs: rhs}
		}
	}

	for {
		tok = p.peek()
		if tok.Type == token.Eof || !tok.Type.IsOp() {
			break
		}
		// TODO add EOF check here

		lBp, rBp, ok := infixBindingPower(tok.Type)
		if !ok {
			break
		}
		if lBp < minBp {
			break
		}

		p.next()
		rhs := p.parseExpression(rBp)
		lhs = &ast.BinaryExpr{Lhs: lhs, Op: *tok, Rhs: rhs}
	}

	return lhs
}

func (p *Parser) expectAndConsumeToken(tok token.TokenType) (token.Token, bool) {
	if p.peek().Type == token.Eof {
		return token.Token{}, false
	}

	if p.peek().Type != tok {
		return token.Token{}, false
	}

	token := p.next()
	return *token, true
}

func (p *Parser) pushSyncStack(syncSet map[token.TokenType]struct{}) {
	p.currentSyncStack = append(p.currentSyncStack, syncSet)
}

func (p *Parser) popSyncStack() {
	if len(p.currentSyncStack) > 0 {
		p.currentSyncStack = p.currentSyncStack[:len(p.currentSyncStack)-1]
	}
}

func prefixBindingPower(op token.TokenType) (struct{}, uint8) {
	switch op {
	case token.Tilde:
		fallthrough
	case token.Plus:
		fallthrough
	case token.Minus:
		return struct{}{}, 13
	}

	return struct{}{}, 0
}

func infixBindingPower(op token.TokenType) (uint8, uint8, bool) {
	switch op {
	case token.BarBar:
		return 1, 2, true
	case token.AmpersandAmpersand:
		return 3, 4, true
	case token.NotEquals:
		fallthrough
	case token.Equals:
		return 5, 6, true
	case token.GreaterEqual:
		fallthrough
	case token.LessEqual:
		fallthrough
	case token.Greater:
		fallthrough
	case token.Less:
		return 7, 8, true
	case token.Plus:
		fallthrough
	case token.Minus:
		return 9, 10, true
	case token.Slash:
		fallthrough
	case token.Percent:
		fallthrough
	case token.Star:
		return 11, 12, true
	}

	return 0, 0, false
}
