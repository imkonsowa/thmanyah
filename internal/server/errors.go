package server

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

func toNetworkError(err error, logger *log.Helper) error {
	if err == nil {
		return nil
	}

	var e *errors.Error
	switch {
	case errors.As(err, &e):
		return err
	default:
		logger.Warn("received unknown error", "error", err)
		
		return errors.InternalServer("INTERNAL_SERVER_ERROR", err.Error())
	}
}
