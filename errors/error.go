package errors

import "errors"

var (
	ErrInvalidAccessToken = errors.New("access token invalid")
	ErrInvalidRespStatus  = errors.New("resp status invalid")
)
