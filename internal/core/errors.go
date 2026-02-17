package core

import "fmt"

type ErrorApp struct {
	Code    string
	Message string
}

func (e ErrorApp) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func Validation(field, msg string) ErrorApp {
	return ErrorApp{
		Code:    "VALIDATION",
		Message: fmt.Sprintf("field '%s': %s", field, msg),
	}
}

func EnvKeyNotFaund(field string) ErrorApp {
	return ErrorApp{
		Code:    "ENV_FILE_LOAD_DATA_ERROR",
		Message: fmt.Sprintf("data: '%s' in .env file not faund. Check your key", field),
	}
}

func JWTMethodError() ErrorApp {
	return ErrorApp{
		Code:    "JWT_METHOD_NOT_VALID",
		Message: "JWT method must be HS256",
	}
}

func JWTTokenNotValid() ErrorApp {
	return ErrorApp{
		Code:    "TOKEN_NOT_VALID",
		Message: "JWT token not valid",
	}
}
