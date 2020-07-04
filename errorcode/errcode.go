package errorcode

type ErrorCode uint32

const (
	OK 						ErrorCode	= 0
	ErrorCodeInvalid 		ErrorCode	= 10001
	ErrorCodeEngineExist 	ErrorCode	= 10002
	ErrorCodeEngineNotExist 	ErrorCode	= 10003
)
