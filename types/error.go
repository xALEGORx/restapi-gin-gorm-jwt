package types

type ApiError struct {
	Code int
	Msg  string
}

var (
	LOGIN_UNKNOWN = &ApiError{202, "user does not exist"}
	LOGIN_ERROR   = &ApiError{203, "wrong login or password"}
	VALID_ERROR   = &ApiError{300, "wrong parameters"}
	ERROR         = &ApiError{400, "operation failed"}
	UNAUTHORIZED  = &ApiError{401, "unauthorized"}
	NOT_FOUND     = &ApiError{404, "not found"}
)

func (e *ApiError) Error() string {
	return e.Msg
}
