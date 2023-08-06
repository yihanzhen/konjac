package errors

import (
	"errors"
	"fmt"
)

var (
	IllegalStateError  = errors.New("Program is in an illegal state")
	IllegalArgError    = errors.New("Input argument is illegal")
	UnimplementedError = errors.New("Not yet implemented")
)

var (
	IllegalAppenderStateError       = fmt.Errorf("%w", IllegalStateError)
	IllegalAppenderLeximeError      = fmt.Errorf("illegal lexime: %w", IllegalAppenderStateError)
	IllegalAppenderConjugationError = fmt.Errorf("illegal conjugation: %w", IllegalAppenderStateError)
)

var (
	UnknownWordError = fmt.Errorf("Unknown word: %w", IllegalArgError)
)
