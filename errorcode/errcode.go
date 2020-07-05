package errorcode

import (
	"encoding/json"
)

type ErrorCode struct {
	Code 	uint32	`json:"code"`
	Message string	`json:"message"`
}

func New(code uint32, message string) ErrorCode{
	return ErrorCode{
		Code:    code,
		Message: message,
	}
}

func (e ErrorCode)ToJson() []byte {
	b,err := json.Marshal(&e)
	if err != nil {
		return []byte("")
	}

	return b
}

var (
	OK 								ErrorCode	= New(0, "")
	ErrorCodeInvalid 				ErrorCode	= New(10001, "")
	ErrorCodeEngineExist 			ErrorCode	= New(10002, "")
	ErrorCodeEngineNotExist 		ErrorCode	= New(10003, "")

	ErrorCodeParamInvalidSymbol		ErrorCode = New(20001, "invalid symbol")
)
