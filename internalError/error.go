package internalError

// IError represents service error
type IError struct {
	ErrorCode uint32
	Message   string
}

func Error(code uint32, msg string) *IError {
	return &IError{
		ErrorCode: code,
		Message:   msg,
	}
}

const (
	DatabaseError  = 50301
	ScanError      = 50302
	HashError      = 42201
	UniqueKeyError = 40001
	NotFoundError  = 40401
	AuthError      = 40301
	JwtError       = 50002
)
