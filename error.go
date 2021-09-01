package err2

type Error struct {
	err              error
	clientMessage    string
	clientStatusCode int
}

func NewClientErr(parentErr error, clientMessage string, clientStatusCode int) Error {
	return Error{
		err:              parentErr,
		clientMessage:    clientMessage,
		clientStatusCode: clientStatusCode,
	}
}

func (e Error) Log(parent, child string) Error {
	LogErr1(parent, child, e)
	return e
}

func NewServerErr(parentErr error) Error {
	return Error{
		err: parentErr,
	}
}

func (e Error) ClientReadable() bool {
	if e.clientMessage == "" {
		return false
	}
	return true
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) ClientMessage() string {
	return e.clientMessage
}

func (e Error) ClientStatusCode() int {
	return e.clientStatusCode
}

func (e *Error) WrapError(err error) {
	e.err = err
}
