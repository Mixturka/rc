package ast

import (
	"fmt"
	"strings"

	"github.com/Mixturka/rc/internal/lexer/token"
)

type ScopableNode interface {
	ScopeStart() int
	ScopeEnd() int
}

type PrintableNode interface {
	Print(src string, sb *strings.Builder, nestingLevel int)
}

type Node interface {
	PrintableNode
	ScopableNode
}

type Stmt interface {
	Node
}

type Expr interface {
	Node
}

type DummyNode struct {
	token token.Token
}

type Program struct {
	Function Func
}

type Func struct {
	Name token.Token
	Body Stmt
}

type ReturnStmt struct {
	Expr Expr
}

type ConstExpr struct {
	Value token.Token
}

func (d *DummyNode) Print(src string, sb *strings.Builder, nestingLevel int) {
	writeIndent(sb, nestingLevel)
	sb.WriteString("Unresolved\n")
}

func (d *DummyNode) ScopeStart() int {
	return d.token.Scope.Start
}

func (d *DummyNode) ScopeEnd() int {
	return d.token.Scope.Start
}

func (f *Func) Print(src string, sb *strings.Builder, nestingLevel int) {
	writeIndent(sb, nestingLevel)
	sb.WriteString("func: {\n")

	nestingLevel++
	writeIndent(sb, nestingLevel)
	fmt.Fprintf(sb, "Name: %s\n", src[f.Name.Scope.Start:f.Name.Scope.End+1])

	writeIndent(sb, nestingLevel)
	sb.WriteString("Body: {\n")
	nestingLevel++

	f.Body.Print(src, sb, nestingLevel)
	nestingLevel--

	writeIndent(sb, nestingLevel)
	sb.WriteString("}\n")
	nestingLevel--

	writeIndent(sb, nestingLevel)
	sb.WriteString("}\n")
}

func (f *Func) ScopeStart() int {
	return f.Name.Scope.Start - 2 // -2 is for 'fn'
}

func (f *Func) ScopeEnd() int {
	return f.Body.ScopeEnd() + 1 // +1 is for '}'
}

func (rs *ReturnStmt) Print(src string, sb *strings.Builder, nestingLevel int) {
	writeIndent(sb, nestingLevel)
	sb.WriteString("return ")

	rs.Expr.Print(src, sb, nestingLevel)
	sb.WriteString(";\n")
}

func (rs *ReturnStmt) ScopeStart() int {
	return rs.Expr.ScopeStart()
}

func (rs *ReturnStmt) ScopeEnd() int {
	return rs.Expr.ScopeEnd()
}

func (ce *ConstExpr) Print(src string, sb *strings.Builder, nestingLevel int) {
	sb.WriteString(src[ce.Value.Scope.Start : ce.Value.Scope.End+1])
}

func (ce *ConstExpr) ScopeStart() int {
	return ce.Value.Scope.Start
}

func (ce *ConstExpr) ScopeEnd() int {
	return ce.Value.Scope.End
}

func writeIndent(sb *strings.Builder, nestingLevel int) {
	for range nestingLevel {
		sb.WriteString("  ")
	}
}
