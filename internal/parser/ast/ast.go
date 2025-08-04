package ast

import "github.com/Mixturka/rc/internal/pkg/scope"

type Program struct {
	Function Function
}

type Function struct {
	Scope scope.Scope
}

type Block struct {
}

type Statement struct {
}

type Expression struct {
}
