package internalError

import "fmt"

// IError represents service error
type IError struct {
	ErrorCode uint32
	Message   string
}

func (i IError) Error() string {
	return fmt.Sprintf("status %d: err %v", i.ErrorCode, i.Message)
}

func MakeError(code uint32, msg string) *IError {
	return &IError{
		ErrorCode: code,
		Message:   msg,
	}
}

const (
	ServerError   = 50001
	JwtError      = 50002
	DatabaseError = 50301
	ScanError     = 50302

	HashError          = 42201
	UniqueKeyError     = 40001
	NotFoundError      = 40401
	Forbidden          = 40301
	InvalidAccessToken = 40302
	AuthError          = 40303
)
