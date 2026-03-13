package errors

import "net/http"

type Kind string

const (
  KindValidation Kind = "validation"
  KindNotFound   Kind = "not_found"
  KindConflict   Kind = "conflict"
  KindForbidden  Kind = "forbidden"
  KindUnauthorized Kind = "unauthorized"
  KindInternal   Kind = "internal"
)

type AppError struct {
  Kind    Kind
  Message string
  Details map[string]string
  Err     error
}

func (e AppError) Error() string {
  if e.Message != "" {
    return e.Message
  }
  if e.Err != nil {
    return e.Err.Error()
  }
  return "error"
}

func HTTPStatus(err error) int {
  ae, ok := err.(AppError)
  if !ok {
    return http.StatusInternalServerError
  }
  switch ae.Kind {
  case KindValidation:
    return http.StatusBadRequest
  case KindNotFound:
    return http.StatusNotFound
  case KindConflict:
    return http.StatusConflict
  case KindForbidden:
    return http.StatusForbidden
  case KindUnauthorized:
    return http.StatusUnauthorized
  default:
    return http.StatusInternalServerError
  }
}