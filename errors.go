package deadline

import "fmt"

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Error int

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	ErrSuccess Error = iota
	ErrBadParameter
	ErrUnexpectedResponse
	ErrInternalAppError
	ErrNotFound
)

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (e Error) Error() string {
	switch e {
	case ErrSuccess:
		return "ErrSuccess"
	case ErrBadParameter:
		return "ErrBadParameter"
	case ErrUnexpectedResponse:
		return "ErrUnexpectedResponse"
	case ErrInternalAppError:
		return "ErrInternalAppError"
	case ErrNotFound:
		return "ErrNotFound"
	default:
		return "[?? Invalid Error value]"
	}
}

// With appends any additional arguments onto an error for context
func (e Error) With(args ...interface{}) error {
	return fmt.Errorf("%w: %v", e, fmt.Sprint(args...))
}
