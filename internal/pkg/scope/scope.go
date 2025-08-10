package scope

// Start and End fields are counted as number of runes from source beginning
type Scope struct {
	Start int
	End   int
	Line  int
}
