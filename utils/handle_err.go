package utils

type HandlerErr struct {
	Code    int
	Message string
	Detail  []ErrorResponse
}

func (h HandlerErr) Error() string {
	return h.Message
}
