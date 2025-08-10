package erremitter

import "errors"

const (
	maxErrors = 20
)

type ErrScope struct {
	Start int
	End   int
}

type SquiggleScope struct {
	Start int
	End   int
	Lines int
}

type ErrType int32

var (
	ErrMaxReached = errors.New("maximum number of errors reached")
)

type Err struct {
	Message   string
	ErrScope  ErrScope
	Squiggles []SquiggleScope
}

type ErrEmitter struct {
	errors []Err
}

func NewErrEmitter() ErrEmitter {
	return ErrEmitter{
		errors: make([]Err, 0, maxErrors),
	}
}

func (ee *ErrEmitter) AddErr(message string, errScope ErrScope, squiggleScopes []SquiggleScope) error {
	if len(ee.errors) >= maxErrors {
		return ErrMaxReached
	}
	ee.errors = append(ee.errors, Err{message, errScope, squiggleScopes})

	return nil
}

func (ee *ErrEmitter) Errors() []Err {
	return ee.errors
}
