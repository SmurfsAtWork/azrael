package errors

import "errors"

var (
	ErrInvalidToken      = errors.New("invalid-session-token")
	ErrExpiredToken      = errors.New("expired-session-token")
	ErrInvalidCredenials = errors.New("invalid-login-credentials")

	ErrSomethingWentWrong = errors.New("something-went-wrong")
)

type PapaError interface {
	error
	Id() string
}

func IsPapa(err error) bool {
	var papaError PapaError
	return errors.As(err, &papaError)
}

func Is(err, target error) bool {
	if IsPapa(err) {
		return true
	}

	return errors.Is(err, target)
}
