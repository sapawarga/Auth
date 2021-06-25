package error

import "errors"

var (
	ErrorIncomingContext = errors.New("error_incoming_context")
	ErrorGetContext      = errors.New("error_get_context_value")
)
