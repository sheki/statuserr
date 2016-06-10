package statuserr

import "fmt"

type err struct {
	StatusCode int
	Message    string
}

// New returns a new status err with a status code
func New(code int, message string) error {
	return err{StatusCode: code, Message: message}
}

// Wrap an existing error with the status code
//  BadRequest.Wrap(errors.New("bad user request"))
func (e err) Wrap(new error) error {
	return err{StatusCode: e.Status(), Message: new.Error()}
}

func (e err) Error() string {
	return e.Message
}

// Status returns the status code associated with the error
func (e err) Status() int {
	return e.StatusCode
}

func (e err) Msg(m string) error {
	return err{StatusCode: e.Status(), Message: m}
}

func (e err) Msgf(format string, a ...interface{}) error {
	return err{StatusCode: e.Status(), Message: fmt.Sprintf(format, a...)}
}
