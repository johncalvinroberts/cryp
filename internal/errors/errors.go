package errors

import (
	goerrs "errors"
)

var (
	ErrValidationFailure = goerrs.New("validation failed")

	ErrDataCreationFailure = goerrs.New("data creation failure")
	ErrDataAccessFailure   = goerrs.New("data access failure")
	ErrDataUpdateFailure   = goerrs.New("data update failure")
	ErrDataDeletionFailure = goerrs.New("data deletion failure")
	ErrInternalServerError = goerrs.New("something unexpected went wrong")
)
