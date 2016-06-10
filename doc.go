// Package statusErr provides a Errors with a HTTP status.
// this is useful for writing HTTP APIs where you want to return a correct
// http status from any part of the codebase
//   _, Err := ioutil.ReadAll(r)
//   if Err != nil {
//     return statusErr.InternalServerError.Wrap(Err)
//   }
//
// other possible uses
//   if user != "stark" {
//     return statusErr.NotImplemented.Msgf("service does not work for %s", user)
//   }
//
// if one is lazy
//   if user.NoName() {
//     return statusErr.Conflict
//   }
// package has types for all http status codes greater than 400.
// The Error values implement interface
//   type StatusCode interface {
//     Status() int
//   }
// to get the status code do
//   if e, ok := Err.(StatusCode); ok {
//     http.Error(w, e.Status(), e.Error())
//   }

package statuserr
