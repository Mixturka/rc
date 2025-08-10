package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Mixturka/rc/internal/codegen"
	"github.com/Mixturka/rc/internal/erremitter"
	"github.com/Mixturka/rc/internal/lexer"
	"github.com/Mixturka/rc/internal/parser"
)

func main() {
	src := "fn main() -> i32 { return 23; }"
	l := lexer.NewLexer([]rune(src))
	tokens, err := l.Tokenize()

	if err != nil {
		log.Fatal("failed to tokenize source")
	}
	em := erremitter.NewErrEmitter()
	p := parser.NewParser(tokens, &em)

	ast := p.Parse()
	var sb strings.Builder
	// ast.Function.Print(src, &sb, 0)

	// fmt.Println(sb.String())
	// fmt.Println(em.Errors())
	cg := codegen.NewCodeGenerator(&sb, src)
	cg.EmitProgram(*ast)

	fmt.Println(sb.String())
}
