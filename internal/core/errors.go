package core

import (
	"errors"
	"fmt"
)

type ErrorApp struct {
	Code    string
	Message string
}

func (e ErrorApp) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func IsError(err error, code string) bool {
	var appErr ErrorApp

	if errors.As(err, &appErr) {
		return appErr.Code == code
	}

	return false
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

func HaveRegister(login string) ErrorApp {
	return ErrorApp{
		Code:    "USER_HAVE_REGISTER",
		Message: fmt.Sprintf("user with login %s had already registered", login),
	}
}

func InvalidPassword() ErrorApp {
	return ErrorApp{
		Code:    "INVALID_PASSWORD",
		Message: "invalid password",
	}
}

func TooShortPassword() ErrorApp {
	return ErrorApp{
		Code:    "PASSWORD_IS_SHORT",
		Message: "The password is too short",
	}
}

func TooLongPassword() ErrorApp {
	return ErrorApp{
		Code:    "PASSWORD_IS_lONG",
		Message: "The password is too long",
	}
}
