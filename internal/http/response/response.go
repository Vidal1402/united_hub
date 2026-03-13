package response

import (
  "encoding/json"
  "net/http"
)

type ErrorResponse struct {
  Error   string            `json:"error"`
  Details map[string]string `json:"details,omitempty"`
}

func JSON(w http.ResponseWriter, status int, v any) {
  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  w.WriteHeader(status)
  if v == nil {
    return
  }
  _ = json.NewEncoder(w).Encode(v)
}

func Error(w http.ResponseWriter, status int, msg string, details map[string]string) {
  JSON(w, status, ErrorResponse{Error: msg, Details: details})
}