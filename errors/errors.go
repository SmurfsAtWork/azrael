package errors

import "errors"

var (
	ErrInvalidToken      = errors.New("Invalid session token")
	ErrExpiredToken      = errors.New("Expired session token")
	ErrInvalidCredenials = errors.New("Invalid login credentials")

	ErrSomethingWentWrong = errors.New("Something went wrong")
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
