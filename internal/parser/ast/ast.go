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
	Accept(emitter CodeEmitter)
}

type Stmt interface {
	Node
}

type Expr interface {
	Node
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

type UnaryExpr struct {
	Op  token.Token
	Rhs Expr
}

type BinaryExpr struct {
	Lhs Expr
	Op  token.Token
	Rhs Expr
}

type ConstExpr struct {
	Value token.Token
}

func (p Program) Accept(emitter CodeEmitter) {
	emitter.EmitProgram(p)
}

func (f Func) Accept(emitter CodeEmitter) {
	emitter.EmitFunc(f)
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

func (rs ReturnStmt) Accept(emitter CodeEmitter) {
	emitter.EmitReturnStmt(rs)
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

func (ux UnaryExpr) Accept(emitter CodeEmitter) {
	emitter.EmitUnaryExpr(ux)
}

func (ux *UnaryExpr) Print(src string, sb *strings.Builder, nestingLevel int) {
	sb.WriteString(src[ux.Op.Scope.Start : ux.Op.Scope.End+1])
	ux.Rhs.Print(src, sb, nestingLevel)
}

func (ux *UnaryExpr) ScopeStart() int {
	return ux.Op.Scope.Start
}

func (ux *UnaryExpr) ScopeEnd() int {
	return ux.Rhs.ScopeEnd()
}

func (bx BinaryExpr) Accept(emitter CodeEmitter) {
	emitter.EmitBinaryExpr(bx)
}

func (bx *BinaryExpr) Print(src string, sb *strings.Builder, nestingLevel int) {
	bx.Lhs.Print(src, sb, nestingLevel)
	sb.WriteString(src[bx.Op.Scope.Start : bx.Op.Scope.End+1])
	bx.Rhs.Print(src, sb, nestingLevel)
}

func (bx *BinaryExpr) ScopeStart() int {
	return bx.Lhs.ScopeStart()
}

func (bx *BinaryExpr) ScopeEnd() int {
	return bx.Rhs.ScopeEnd()
}

func (ce ConstExpr) Accept(emitter CodeEmitter) {
	emitter.EmitConstExpr(ce)
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

type CodeEmitter interface {
	EmitProgram(program Program)
	EmitFunc(fn Func)
	EmitUnaryExpr(expr UnaryExpr)
	EmitBinaryExpr(expr BinaryExpr)
	EmitReturnStmt(stmt ReturnStmt)
	EmitConstExpr(expr ConstExpr)
}
