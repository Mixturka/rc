package parser

import (
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
}

func NewParser(tokens []token.Token, errEmitter *erremitter.ErrEmitter) Parser {
	return Parser{
		tokens:           tokens,
		errEmitter:       errEmitter,
		currentSyncStack: make([]map[token.TokenType]struct{}, 0),
	}
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

	expr := p.parseExpression()

	if _, ok := p.expectAndConsumeToken(token.Semicolon); !ok {
		log.Fatalf("expected ';' in the end of statement")
	}

	return &ast.ReturnStmt{Expr: expr}
}

func (p *Parser) parseExpression() ast.Expr {
	p.pushSyncStack(exprSyncSet)
	defer p.popSyncStack()

	var constant token.Token
	constant, ok := p.expectAndConsumeToken(token.IntegerNumber)
	if !ok {
		log.Fatalf("expected integer")
	}

	return &ast.ConstExpr{
		Value: constant,
	}
}

func (p *Parser) expectAndConsumeToken(tok token.TokenType) (token.Token, bool) {
	if p.tokens[p.pos].Type == token.Eof {
		return token.Token{}, false
	}

	if p.inErr {
		if len(p.currentSyncStack) <= 0 {
			return token.Token{}, false
		}

		syncSet := p.currentSyncStack[len(p.currentSyncStack)-1]
		if _, ok := syncSet[p.tokens[p.pos].Type]; ok {
			p.inErr = false
		} else {
			p.pos++
		}

		return p.tokens[p.pos-1], true
	}

	if p.tokens[p.pos].Type != tok {
		return token.Token{}, false
	}

	p.pos++
	return p.tokens[p.pos-1], true
}

func (p *Parser) pushSyncStack(syncSet map[token.TokenType]struct{}) {
	p.currentSyncStack = append(p.currentSyncStack, syncSet)
}

func (p *Parser) popSyncStack() {
	if len(p.currentSyncStack) > 0 {
		p.currentSyncStack = p.currentSyncStack[:len(p.currentSyncStack)-1]
	}
}
