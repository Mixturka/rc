package erremitter

import "errors"

const (
	maxErrors = 20
)

type ErrType int32

var (
	ErrMaxReached = errors.New("maximum number of errors reached")
)

type Err struct {
	Message string
	Start   int
	End     int
}

type ErrEmitter struct {
	errors []Err
}

func NewErrEmitter() ErrEmitter {
	return ErrEmitter{
		errors: make([]Err, maxErrors),
	}
}

func (ee *ErrEmitter) AddErr(message string, start, end int) error {
	if len(ee.errors) >= maxErrors {
		return ErrMaxReached
	}
	ee.errors = append(ee.errors, Err{message, start, end})

	return nil
}

func (ee *ErrEmitter) Errors() []Err {
	return ee.errors
}
