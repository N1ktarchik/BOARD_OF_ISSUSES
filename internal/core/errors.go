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

func NullPassword() ErrorApp {
	return ErrorApp{
		Code:    "PASSWORD_IS_NULL",
		Message: "password can not be null",
	}
}

func NullName() ErrorApp {
	return ErrorApp{
		Code:    "NAME_IS_NULL",
		Message: "name can not be null",
	}
}

func NullLogin() ErrorApp {
	return ErrorApp{
		Code:    "LOGIN_IS_NULL",
		Message: "login can not be null",
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
