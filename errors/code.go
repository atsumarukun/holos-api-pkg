package errors

type ErrorCode string

func (ec ErrorCode) String() string {
	return string(ec)
}

var (
	CodeBadRequest          ErrorCode = "BAD_REQUEST"
	CodeUnauthenticated     ErrorCode = "UNAUTHENTICATED"
	CodeUnauthorized        ErrorCode = "UNAUTHORIZED"
	CodeNotFound            ErrorCode = "NOT_FOUND"
	CodeDuplicate           ErrorCode = "DUPLICATE"
	CodeConstraintViolation ErrorCode = "CONSTRAINT_VIOLATION"
	CodeInvalidInput        ErrorCode = "INVALID_INPUT"
	CodeInternalServerError ErrorCode = "INTERNAL_SERVER_ERROR"
	CodeUnknown             ErrorCode = "UNKNOWN"
)
