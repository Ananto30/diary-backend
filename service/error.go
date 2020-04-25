package service

// SError represents service error
type SError struct {
	ErrorCode uint32
	Message   string
}

func Error(code uint32, msg string) *SError {
	return &SError{
		ErrorCode: code,
		Message:   msg,
	}
}

const (
	DatabaseError = 50301
	ScanError     = 50302
	HashError     = 42201
)
