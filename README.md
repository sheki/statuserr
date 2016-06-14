Package statuserr provides a errors with a HTTP status.
this is useful for writing HTTP APIs where you want to return a correct
http status from any part of the codebase

Inspired by [nodejs boom](https://github.com/hapijs/boom)

##Usage

    _, err := ioutil.ReadAll(r)
    if err != nil {
      return statuserr.InternalServerError.Wrap(err)
    }

other possible uses

    if user != "stark" {
      return statuserr.NotImplemented.Msgf("service does not work for %s", user)
    }

if one is lazy

    if user.NoName() {
      return statuserr.Conflict
    }
package has types for all http status codes greater than 400.


The error values implement interface

    type StatusCode interface {
      Status() int
    }

to get the status code do

    if e, ok := err.(StatusCode); ok {
      http.Error(w, e.Status(), e.Error())
    }
