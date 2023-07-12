package errors

import "errors"

var (
	IllegalStateError  = errors.New("Program is in an illegal state")
	IllegalArgError    = errors.New("Input argument is illegal")
	UnimplementedError = errors.New("Not yet implemented")
)
