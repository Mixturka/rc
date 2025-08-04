package scope

type Scope struct {
	Start int
	End   int
}

func NewScope(start, end int) Scope {
	return Scope{start, end}
}
