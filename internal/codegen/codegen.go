package codegen

import (
	"io"
	"strings"

	"github.com/Mixturka/rc/internal/parser/ast"
)

type CodeGenerator struct {
	w     io.Writer
	ident int
	sb    strings.Builder
	src   string
}

func NewCodeGenerator(w io.Writer, src string) CodeGenerator {
	return CodeGenerator{w: w, ident: 0, src: src}
}

func (cg *CodeGenerator) EmitProgram(program ast.Program) {
	program.Function.Accept(cg)
	cg.w.Write([]byte(cg.sb.String()))
}

func (cg *CodeGenerator) EmitFunc(fn ast.Func) {
	cg.writeIndent()
	cg.sb.WriteString("int ")
	cg.sb.WriteString(cg.src[fn.Name.Scope.Start : fn.Name.Scope.End+1])
	cg.sb.WriteString("() {\n")

	cg.ident++
	cg.writeIndent()

	fn.Body.Accept(cg)
	cg.ident--
	cg.writeIndent()
	cg.sb.WriteString("}\n")
}

func (cg *CodeGenerator) EmitUnaryExpr(expr ast.UnaryExpr) {
	cg.sb.WriteRune('(')
	cg.sb.WriteString(cg.src[expr.Op.Scope.Start : expr.Op.Scope.End+1])
	expr.Rhs.Accept(cg)
	cg.sb.WriteRune(')')
}

func (cg *CodeGenerator) EmitBinaryExpr(expr ast.BinaryExpr) {
	cg.sb.WriteRune('(')
	expr.Lhs.Accept(cg)
	cg.sb.WriteRune(' ')
	cg.sb.WriteString(cg.src[expr.Op.Scope.Start : expr.Op.Scope.End+1])
	cg.sb.WriteRune(' ')
	expr.Rhs.Accept(cg)
	cg.sb.WriteRune(')')
}

func (cg *CodeGenerator) EmitReturnStmt(stmt ast.ReturnStmt) {
	cg.writeIndent()
	cg.sb.WriteString("return ")
	stmt.Expr.Accept(cg)
	cg.sb.WriteString(";\n")
}

func (cg *CodeGenerator) EmitConstExpr(expr ast.ConstExpr) {
	cg.sb.WriteString(cg.src[expr.Value.Scope.Start : expr.Value.Scope.End+1])
}

func (cg *CodeGenerator) writeIndent() {
	for range cg.ident {
		cg.sb.WriteString("  ")
	}
}
