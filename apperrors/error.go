package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

var _ error = (*MyAppError)(nil)

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}